package server

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	payment "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	user "github.com/NeptuneG/go-back/gen/go/services/user/proto"
	db "github.com/NeptuneG/go-back/services/payment/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	timeout = time.Second * 5
)

type PaymentService struct {
	payment.UnimplementedPaymentServiceServer
	userClient user.UserServiceClient
	liveClient live.LiveServiceClient
	store      *db.Store
}

func New(ctx context.Context, dbConn *sql.DB) (*PaymentService, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
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
		userClient: user.NewUserServiceClient(userConn),
		liveClient: live.NewLiveServiceClient(liveConn),
		store:      db.NewStore(dbConn),
	}, nil
}

var errs [3]error

func (s *PaymentService) CreateLiveEventOrder(ctx context.Context, req *payment.CreateLiveEventOrderRequest) (*payment.CreateLiveEventOrderResponse, error) {
	if err := s.validateRequest(ctx, req); err != nil {
		return nil, err
	}

	liveEventOrder, err := s.initLiveEventOrder(ctx, req)
	if err != nil {
		log.Printf("failed to init live event order: %v", err)
		return nil, err
	}
	log.Println("live event order created:", liveEventOrder)
	var wg sync.WaitGroup

	if req.UserPoints != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.consumeUserPoints(ctx, req); err != nil {
				log.Printf("failed to consume user points: %v", err)
			}
			log.Println("consumed user points:", req.UserPoints)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		creditCardPayment := req.Price - req.UserPoints
		if err := s.consumeUserCreditCard(ctx, creditCardPayment); err != nil {
			log.Printf("failed to consume user credit card: %v", err)
		}
		log.Println("consumed user credit card:", creditCardPayment)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.reserveSeat(ctx, req.LiveEventId); err != nil {
			log.Printf("failed to reserve a seat: %v", err)
		}
		log.Println("reserved a seat for", req.LiveEventId)
	}()

	wg.Wait()
	if err := s.completeliveEventOrder(ctx, liveEventOrder); err != nil {
		log.Printf("failed to complete live event order: %v", err)
	}
	log.Println("complete live event order:", liveEventOrder)

	return &payment.CreateLiveEventOrderResponse{
		Status: "ok",
	}, nil
}

func (s *PaymentService) validateRequest(ctx context.Context, req *payment.CreateLiveEventOrderRequest) error {
	userExistResp, err := s.userClient.IsUserExist(ctx, &user.IsUserExistRequest{
		Id: req.UserId,
	})
	if err != nil {
		return err
	}
	if !userExistResp.Exist {
		return errors.New("user not exist")
	}

	liveEventExistResp, err := s.liveClient.IsLiveEventExist(ctx, &live.IsLiveEventExistRequest{
		Id: req.LiveEventId,
	})
	if err != nil {
		return err
	}
	if !liveEventExistResp.Exist {
		return errors.New("live event not exist")
	}
	return nil
}

func (s *PaymentService) initLiveEventOrder(ctx context.Context, req *payment.CreateLiveEventOrderRequest) (*db.LiveEventOrder, error) {
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

func (s *PaymentService) consumeUserPoints(ctx context.Context, req *payment.CreateLiveEventOrderRequest) error {
	done := make(chan error, 1)
	go func() {
		var err error
		for {
			_, err := s.userClient.ConsumeUserPoints(ctx, &user.ConsumeUserPointsRequest{
				Id:          req.UserId,
				Points:      req.UserPoints,
				Description: "consume user points for live event order",
			})
			if err == nil {
				break
			}
		}
		done <- err
	}()
	select {
	case <-time.After(time.Duration(timeout)):
		errs[0] = errors.New("timeout")
	case errs[0] = <-done:
	}
	return errs[0]
}

func (s *PaymentService) consumeUserCreditCard(ctx context.Context, payment int32) error {
	done := make(chan error, 1)
	go func() {
		var err error
		for {
			result := rand.Intn(10)
			if result%10 == 5 {
				done <- errors.New("credit card error")
			} else {
				done <- err
			}
			time.Sleep(time.Second)
		}
	}()
	select {
	case <-time.After(time.Duration(timeout)):
		log.Println("credit timout")
		errs[1] = errors.New("timeout")
	case errs[1] = <-done:
	}
	return errs[1]
}

func (s *PaymentService) reserveSeat(ctx context.Context, liveEventID string) error {
	done := make(chan error, 1)
	go func() {
		var err error
		for {
			_, err := s.liveClient.ReserveSeat(ctx, &live.ReserveSeatRequest{
				LiveEventId: liveEventID,
			})
			if err == nil {
				break
			}
		}
		done <- err
	}()
	select {
	case <-time.After(time.Duration(timeout)):
		errs[2] = errors.New("timeout")
	case errs[2] = <-done:
	}
	return errs[2]
}

func (s *PaymentService) completeliveEventOrder(ctx context.Context, liveEventOrder *db.LiveEventOrder) error {
	log.Println("begin completeliveEventOrder")
	isAllDone := true
	for _, err := range errs {
		if err != nil {
			log.Println("isAllDone = false")
			isAllDone = false
		}
	}
	var state db.State
	if isAllDone {
		state = db.StatePaid
	} else {
		state = db.StateFailed
	}
	return s.store.UpdateLiveEventOrderState(ctx, db.UpdateLiveEventOrderStateParams{
		ID:    liveEventOrder.ID,
		State: state,
	})
}
