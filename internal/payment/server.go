package payment

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	liveProto "github.com/NeptuneG/go-back/api/proto/live"
	paymentProto "github.com/NeptuneG/go-back/api/proto/payment"
	db "github.com/NeptuneG/go-back/internal/payment/db/sqlc"
	"github.com/NeptuneG/go-back/internal/pkg/db/types"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	timeout = time.Second * 5
)

//go:embed config/retry_policy.json
var retryPolicy string

type PaymentService struct {
	paymentProto.UnimplementedPaymentServiceServer
	liveClient liveProto.LiveServiceClient
	store      *db.Store
}

func New() *PaymentService {
	ctx := context.Background()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}

	liveConn, err := grpc.DialContext(ctx, os.Getenv("LIVE_SERVICE_HOST")+":"+os.Getenv("LIVE_SERVICE_PORT"), opts...)
	if err != nil {
		log.Fatal("failed to connect to live service", log.Field.Error(err))
		panic(err)
	}

	return &PaymentService{
		liveClient: liveProto.NewLiveServiceClient(liveConn),
		store:      db.NewStore(),
	}
}

func (s *PaymentService) Close() {
	if err := s.store.Close(); err != nil {
		log.Fatal("failed to close database connection", log.Field.Error(err))
		panic(err)
	}
}

func (s *PaymentService) CreateUserPoints(ctx context.Context, req *paymentProto.CreateUserPointsRequest) (*paymentProto.CreateUserPointsResponse, error) {
	userID := uuid.MustParse(req.UserId)
	_, err := s.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      userID,
		Points:      req.UserPoints,
		Description: types.NewNullString(fmt.Sprintf("Add %d points", req.UserPoints)),
		OrderType:   "UserPointsOrder",
		OrderID:     uuid.New(),
	})
	if err != nil {
		log.Error("failed to create user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	userPoints, err := s.store.GetUserPoints(ctx, userID)
	if err != nil {
		log.Error("failed to get user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentProto.CreateUserPointsResponse{
		UserId:     req.UserId,
		UserPoints: userPoints.(int64),
	}, nil
}

func (s *PaymentService) GetUserPoints(ctx context.Context, req *paymentProto.GetUserPointsRequest) (*paymentProto.GetUserPointsResponse, error) {
	userID := uuid.MustParse(req.UserId)
	userPoints, err := s.store.GetUserPoints(ctx, userID)
	if err != nil {
		log.Error("failed to get user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentProto.GetUserPointsResponse{
		UserId:     req.UserId,
		UserPoints: userPoints.(int64),
	}, nil
}

func (s *PaymentService) CreateLiveEventOrder(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) (*paymentProto.CreateLiveEventOrderResponse, error) {
	liveEvent := s.getLiveEventOrderRelations(ctx, req)
	if liveEvent == nil {
		return nil, status.Error(codes.InvalidArgument, "failed to get live event or user")
	}
	if err := s.validateCreateLiveEventOrderRequest(ctx, liveEvent, req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	liveEventOrder, err := s.initLiveEventOrder(ctx, req)
	if err != nil {
		log.Error("failed to init live event order", log.Field.Error(err))
		return nil, status.Error(codes.Internal, "failed to init live event order")
	}
	log.Info("live event order created", log.Field.String("id", liveEventOrder.ID.String()), log.Field.Any("state", liveEventOrder.State))

	var wg sync.WaitGroup
	isSuccess := true

	if req.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.consumeUserPoints(ctx, req, liveEvent, liveEventOrder.ID); err != nil {
				log.Error("failed to consume user points", log.Field.Error(err))
				isSuccess = false
			} else {
				log.Info("user points consumed", log.Field.String("user_id", req.UserId), log.Field.Int32("user_points", req.UserPoints))
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		creditCardPayment := req.Price - req.UserPoints
		if err := s.consumeUserCreditCard(ctx, creditCardPayment); err != nil {
			log.Error("failed to consume user credit card", log.Field.Error(err))
			isSuccess = false
		} else {
			log.Info("user credit card consumed", log.Field.String("user_id", req.UserId), log.Field.Int32("credit_card_payment", creditCardPayment))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.reserveSeat(ctx, req.LiveEventId); err != nil {
			log.Error("failed to reserve seat", log.Field.Error(err))
			isSuccess = false
		} else {
			log.Info("reserved seat", log.Field.String("live_event_id", req.LiveEventId))
		}
	}()

	wg.Wait()
	if err := s.completeliveEventOrder(ctx, liveEventOrder, isSuccess); err != nil {
		log.Error("failed to complete live event order", log.Field.Error(err))
	}
	log.Info("live event order completed", log.Field.String("id", liveEventOrder.ID.String()), log.Field.Any("state", liveEventOrder.State))

	if liveEventOrder.State == db.StateFailed {
		s.rollbackLiveEventOrder(ctx, liveEventOrder)
	}

	return &paymentProto.CreateLiveEventOrderResponse{
		State: string(liveEventOrder.State),
	}, nil
}

func (s *PaymentService) getLiveEventOrderRelations(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) *liveProto.LiveEvent {
	liveEventResp, err := s.liveClient.GetLiveEvent(ctx, &liveProto.GetLiveEventRequest{
		Id: req.LiveEventId,
	})
	if err != nil {
		log.Error("failed to get live event", log.Field.Error(err))
		return nil
	}
	return liveEventResp.LiveEvent
}

func (s *PaymentService) validateCreateLiveEventOrderRequest(ctx context.Context, liveEvent *liveProto.LiveEvent, req *paymentProto.CreateLiveEventOrderRequest) error {
	userID := uuid.MustParse(req.UserId)
	userPoints, err := s.store.GetUserPoints(ctx, userID)
	if err != nil {
		return err
	}
	if userPoints.(int64) < int64(req.UserPoints) {
		return errors.New("user points not enough")
	}
	if liveEvent.AvailableSeats == 0 {
		return errors.New("no seats available")
	}
	if liveEvent.StageOneStartAt.AsTime().Before(time.Now()) {
		return errors.New("live event has started")
	}
	return nil
}

func (s *PaymentService) initLiveEventOrder(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) (*db.LiveEventOrder, error) {
	userID := uuid.MustParse(req.UserId)
	liveEventID, err := uuid.Parse(req.LiveEventId)
	if err != nil {
		return nil, err
	}
	liveEventOrder, err := s.store.CreateLiveEventOrder(ctx, db.CreateLiveEventOrderParams{
		UserID:      userID,
		LiveEventID: liveEventID,
		Price:       req.Price,
		UserPoints:  req.UserPoints,
	})
	return &liveEventOrder, err
}

func (s *PaymentService) consumeUserPoints(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest, liveEvent *liveProto.LiveEvent, orderId uuid.UUID) error {
	done := make(chan error)
	go func() {
		userID := uuid.MustParse(req.UserId)
		_, err := s.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
			UserID:      userID,
			Points:      req.UserPoints,
			Description: types.NewNullString("order: " + liveEvent.Title),
			OrderType:   "LiveEventOrder",
			OrderID:     orderId,
		})
		log.Debug("ConsumeUserPoints responded", log.Field.Any("code", status.Code(err)))
		done <- err
	}()

	select {
	case <-time.After(time.Duration(timeout)):
		return errors.New("timeout")
	case err := <-done:
		return err
	}
}

func (s *PaymentService) consumeUserCreditCard(ctx context.Context, payment int32) error {
	done := make(chan error)
	go func() {
		result := rand.Intn(10)
		// result := 5
		log.Debug("consumeUserCreditCard result", log.Field.Int("result", result))
		if result%10 == 5 {
			done <- errors.New("credit card error")
		} else {
			done <- nil
		}
	}()

	select {
	case <-time.After(time.Duration(timeout)):
		return errors.New("timeout")
	case err := <-done:
		return err
	}
}

func (s *PaymentService) reserveSeat(ctx context.Context, liveEventID string) error {
	done := make(chan error)
	go func() {
		_, err := s.liveClient.ReserveSeat(ctx, &liveProto.ReserveSeatRequest{
			LiveEventId: liveEventID,
		})
		log.Debug("ReserveSeat responded", log.Field.Any("code", status.Code(err)))
		done <- err
	}()

	select {
	case <-time.After(time.Duration(timeout)):
		return errors.New("timeout")
	case err := <-done:
		return err
	}
}

func (s *PaymentService) completeliveEventOrder(ctx context.Context, liveEventOrder *db.LiveEventOrder, isSuccess bool) error {
	log.Info("begin completeliveEventOrder", log.Field.Bool("isSuccess", isSuccess))
	if isSuccess {
		liveEventOrder.State = db.StatePaid
	} else {
		liveEventOrder.State = db.StateFailed
	}
	return s.store.UpdateLiveEventOrderState(ctx, db.UpdateLiveEventOrderStateParams{
		ID:    liveEventOrder.ID,
		State: liveEventOrder.State,
	})
}

func (s *PaymentService) rollbackLiveEventOrder(ctx context.Context, liveEventOrder *db.LiveEventOrder) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := s.liveClient.RollbackSeatReservation(ctx, &liveProto.RollbackSeatReservationRequest{
			LiveEventId: liveEventOrder.LiveEventID.String(),
		})
		log.Debug("RollbackSeatReservation responded", log.Field.Any("code", status.Code(err)))
		if err != nil {
			log.Fatal("failed to rollback seat reservation", log.Field.Error(err))
			// notify for manual follow up
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.store.DeleteUserPointsByOrderID(ctx, liveEventOrder.ID)
		log.Debug("DeleteUserPointsByOrderID responded", log.Field.Any("code", status.Code(err)))
		if err != nil {
			log.Fatal("failed to rollback consume user points", log.Field.Error(err))
			// notify for manual follow up
		}
	}()

	if liveEventOrder.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := rand.Intn(10)
			log.Debug("rollbackConsumeUserCreditCard result", log.Field.Int("result", result))
			if result%10 == 5 {
				log.Fatal("failed to rollback consume user credit card")
			}
		}()
	}
	wg.Wait()
}
