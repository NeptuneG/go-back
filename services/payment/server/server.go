package server

import (
	"context"
	"database/sql"

	"github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	db "github.com/NeptuneG/go-back/services/payment/db/sqlc"
	"github.com/google/uuid"
)

type PaymentService struct {
	proto.UnimplementedPaymentServiceServer
	store *db.Store
}

func New(dbConn *sql.DB) *PaymentService {
	return &PaymentService{
		store: db.NewStore(dbConn),
	}
}

func (liveService *PaymentService) CreateLiveEventOrder(ctx context.Context, req *proto.CreateLiveEventOrderRequest) (*proto.CreateLiveEventOrderResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	liveEventID, err := uuid.Parse(req.LiveEventId)
	if err != nil {
		return nil, err
	}
	_, err = liveService.store.CreateLiveEventOrder(ctx, db.CreateLiveEventOrderParams{
		UserID:      userID,
		LiveEventID: liveEventID,
		Price:       req.Price,
		UserPoints:  req.UserPoints,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateLiveEventOrderResponse{
		Status: "ok",
	}, nil
}
