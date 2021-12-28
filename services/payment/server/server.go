package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"

	liveProto "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	paymentProto "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	userProto "github.com/NeptuneG/go-back/gen/go/services/user/proto"
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
			"name": [{
					"service": "neptuneg.go_back.serivces.live.LiveService",
					"method": "ReserveSeat"
				},
				{
					"service": "neptuneg.go_back.serivces.user.UserService",
					"method": "ConsumeUserPoints"
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
		log.Fatalln("failed to dial user service:", err)
		return nil, err
	}

	liveConn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		log.Fatalln("failed to dial live service:", err)
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
		log.Printf("failed to init live event order: %v", err)
		return nil, err
	}
	log.Println("live event order initialized:", liveEventOrder)

	var wg sync.WaitGroup
	isSuccess := true

	if req.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.consumeUserPoints(ctx, req, liveEvent); err != nil {
				log.Printf("failed to consume user points: %v", err)
				isSuccess = false
			} else {
				log.Println("consumed user points:", req.UserPoints)
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		creditCardPayment := req.Price - req.UserPoints
		if err := s.consumeUserCreditCard(ctx, creditCardPayment); err != nil {
			log.Printf("failed to consume user credit card: %v", err)
			isSuccess = false
		} else {
			log.Println("consumed user credit card:", creditCardPayment)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.reserveSeat(ctx, req.LiveEventId); err != nil {
			log.Printf("failed to reserve a seat: %v", err)
			isSuccess = false
		} else {
			log.Println("reserved a seat for", req.LiveEventId)
		}
	}()

	wg.Wait()
	if err := s.completeliveEventOrder(ctx, liveEventOrder, isSuccess); err != nil {
		log.Printf("failed to complete live event order: %v", err)
	}
	log.Println("complete live event order:", liveEventOrder)

	if liveEventOrder.State == db.StateFailed {
		if err := s.rollbackLiveEventOrder(ctx, liveEventOrder); err != nil {
			log.Printf("failed to rollback live event order: %v", err)
			return nil, err
		}
	}

	return &paymentProto.CreateLiveEventOrderResponse{
		Status: string(liveEventOrder.State),
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
			log.Printf("failed to get live event: %v", err)
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
			log.Printf("failed to get user: %v", err)
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

func (s *PaymentService) consumeUserPoints(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest, liveEvent *liveProto.LiveEvent) error {
	done := make(chan error)
	go func() {
		_, err := s.userClient.ConsumeUserPoints(ctx, &userProto.ConsumeUserPointsRequest{
			Id:          req.UserId,
			Points:      req.UserPoints,
			Description: "order: " + liveEvent.Title,
		})
		log.Println("ConsumeUserPoints responded =", status.Code(err))
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
		time.Sleep(3 * time.Second)
		result := rand.Intn(10)
		// result := 5
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
		log.Println("ReserveSeat responded =", status.Code(err))
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
	log.Printf("begin completeliveEventOrder, isSuccess: %v", isSuccess)
	var state db.State
	if isSuccess {
		state = db.StatePaid
	} else {
		state = db.StateFailed
	}
	return s.store.UpdateLiveEventOrderState(ctx, db.UpdateLiveEventOrderStateParams{
		ID:    liveEventOrder.ID,
		State: state,
	})
}

func (s *PaymentService) rollbackLiveEventOrder(ctx context.Context, liveEventOrder *db.LiveEventOrder) error {
	return nil
}
