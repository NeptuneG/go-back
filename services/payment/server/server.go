package server

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"sync"
	"time"

	liveProto "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	paymentProto "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	userProto "github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	db "github.com/NeptuneG/go-back/services/payment/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	timeout     = time.Second * 5
	retryPolicy = `{
		"methodConfig": [{
			"name": [
				{
					"service": "neptuneg.go_back.serivces.live.LiveService",
					"method": "ReserveSeat"
				},
				{
					"service": "neptuneg.go_back.serivces.live.LiveService",
					"method": "RollbackSeatReservation"
				},
				{
					"service": "neptuneg.go_back.serivces.user.UserService",
					"method": "ConsumeUserPoints"
				},
				{
					"service": "neptuneg.go_back.serivces.user.UserService",
					"method": "RollbackConsumeUserPoints"
				}
			],
			"retryPolicy": {
				"MaxAttempts": 4,
				"InitialBackoff": ".01s",
				"MaxBackoff": ".01s",
				"BackoffMultiplier": 1.0,
				"RetryableStatusCodes": ["INTERNAL"]
			}
		}]
	}`
)

type PaymentService struct {
	paymentProto.UnimplementedPaymentServiceServer
	userClient userProto.UserServiceClient
	liveClient liveProto.LiveServiceClient
	store      *db.Store
}

func New(ctx context.Context, dbConn *sql.DB) (*PaymentService, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}

	userConn, err := grpc.DialContext(ctx, "user-service:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to user service", logField.Error(err))
		return nil, err
	}

	liveConn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		log.Fatal("failed to connect to live service", logField.Error(err))
		return nil, err
	}

	return &PaymentService{
		userClient: userProto.NewUserServiceClient(userConn),
		liveClient: liveProto.NewLiveServiceClient(liveConn),
		store:      db.NewStore(dbConn),
	}, nil
}

func (s *PaymentService) CreateLiveEventOrder(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) (*paymentProto.CreateLiveEventOrderResponse, error) {
	liveEvent, user := s.getLiveEventOrderRelations(ctx, req)
	if liveEvent == nil || user == nil {
		return nil, status.Error(codes.InvalidArgument, "failed to get live event or user")
	}
	if err := s.validateCreateLiveEventOrderRequest(liveEvent, user, req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	liveEventOrder, err := s.initLiveEventOrder(ctx, req)
	if err != nil {
		log.Error("failed to init live event order", logField.Error(err))
		return nil, status.Error(codes.Internal, "failed to init live event order")
	}
	log.Info("live event order created", logField.String("id", liveEventOrder.ID.String()), logField.Any("state", liveEventOrder.State))

	var wg sync.WaitGroup
	isSuccess := true

	if req.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.consumeUserPoints(ctx, req, liveEvent, liveEventOrder.ID); err != nil {
				log.Error("failed to consume user points", logField.Error(err))
				isSuccess = false
			} else {
				log.Info("user points consumed", logField.String("user_id", req.UserId), logField.Int32("user_points", req.UserPoints))
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		creditCardPayment := req.Price - req.UserPoints
		if err := s.consumeUserCreditCard(ctx, creditCardPayment); err != nil {
			log.Error("failed to consume user credit card", logField.Error(err))
			isSuccess = false
		} else {
			log.Info("user credit card consumed", logField.String("user_id", user.Id), logField.Int32("credit_card_payment", creditCardPayment))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.reserveSeat(ctx, req.LiveEventId); err != nil {
			log.Error("failed to reserve seat", logField.Error(err))
			isSuccess = false
		} else {
			log.Info("reserved seat", logField.String("live_event_id", req.LiveEventId))
		}
	}()

	wg.Wait()
	if err := s.completeliveEventOrder(ctx, liveEventOrder, isSuccess); err != nil {
		log.Error("failed to complete live event order", logField.Error(err))
	}
	log.Info("live event order completed", logField.String("id", liveEventOrder.ID.String()), logField.Any("state", liveEventOrder.State))

	if liveEventOrder.State == db.StateFailed {
		s.rollbackLiveEventOrder(ctx, liveEventOrder)
	}

	return &paymentProto.CreateLiveEventOrderResponse{
		State: string(liveEventOrder.State),
	}, nil
}

func (s *PaymentService) getLiveEventOrderRelations(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) (*liveProto.LiveEvent, *userProto.User) {
	liveEventCh := make(chan *liveProto.LiveEvent)
	userCh := make(chan *userProto.User)

	go func() {
		liveEventResp, err := s.liveClient.GetLiveEvent(ctx, &liveProto.GetLiveEventRequest{
			Id: req.LiveEventId,
		})
		if err != nil {
			log.Error("failed to get live event", logField.Error(err))
			liveEventCh <- nil
		} else {
			liveEventCh <- liveEventResp.LiveEvent
		}
	}()

	go func() {
		userResp, err := s.userClient.GetUser(ctx, &userProto.GetUserRequest{
			Id: req.UserId,
		})
		if err != nil {
			log.Error("failed to get user", logField.Error(err))
			userCh <- nil
		} else {
			userCh <- userResp.User
		}
	}()

	return <-liveEventCh, <-userCh
}

func (s *PaymentService) validateCreateLiveEventOrderRequest(liveEvent *liveProto.LiveEvent, user *userProto.User, req *paymentProto.CreateLiveEventOrderRequest) error {
	if user.Points < int64(req.UserPoints) {
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
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
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
		_, err := s.userClient.ConsumeUserPoints(ctx, &userProto.ConsumeUserPointsRequest{
			UserId:      req.UserId,
			Points:      req.UserPoints,
			Description: "order: " + liveEvent.Title,
			OrderId:     orderId.String(),
		})
		log.Debug("ConsumeUserPoints responded", logField.Any("code", status.Code(err)))
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
		log.Debug("consumeUserCreditCard result", logField.Int("result", result))
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
		log.Debug("ReserveSeat responded", logField.Any("code", status.Code(err)))
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
	log.Info("begin completeliveEventOrder", logField.Bool("isSuccess", isSuccess))
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
		log.Debug("RollbackSeatReservation responded", logField.Any("code", status.Code(err)))
		if err != nil {
			log.Fatal("failed to rollback seat reservation", logField.Error(err))
			// notify for manual follow up
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := s.userClient.RollbackConsumeUserPoints(ctx, &userProto.RollbackConsumeUserPointsRequest{
			OrderId: liveEventOrder.ID.String(),
		})
		log.Debug("RollbackConsumeUserPoints responded", logField.Any("code", status.Code(err)))
		if err != nil {
			log.Fatal("failed to rollback consume user points", logField.Error(err))
			// notify for manual follow up
		}
	}()

	if liveEventOrder.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := rand.Intn(10)
			log.Debug("rollbackConsumeUserCreditCard result", logField.Int("result", result))
			if result%10 == 5 {
				log.Fatal("failed to rollback consume user credit card")
			}
		}()
	}
	wg.Wait()
}
