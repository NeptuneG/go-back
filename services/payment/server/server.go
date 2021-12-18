package server

import (
	"context"
	"database/sql"
	"errors"
	"log"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	payment "github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	user "github.com/NeptuneG/go-back/gen/go/services/user/proto"
	db "github.com/NeptuneG/go-back/services/payment/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
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

func (s *PaymentService) CreateLiveEventOrder(ctx context.Context, req *payment.CreateLiveEventOrderRequest) (*payment.CreateLiveEventOrderResponse, error) {
	if err := s.validateRequest(ctx, req); err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	liveEventID, err := uuid.Parse(req.LiveEventId)
	if err != nil {
		return nil, err
	}
	_, err = s.store.CreateLiveEventOrder(ctx, db.CreateLiveEventOrderParams{
		UserID:      userID,
		LiveEventID: liveEventID,
		Price:       req.Price,
		UserPoints:  req.UserPoints,
	})
	if err != nil {
		return nil, err
	}

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
