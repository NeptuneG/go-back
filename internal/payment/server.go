package payment

import (
	"context"
	"errors"
	"os"
	"time"

	liveProto "github.com/NeptuneG/go-back/api/proto/live"
	paymentProto "github.com/NeptuneG/go-back/api/proto/payment"
	liveSvc "github.com/NeptuneG/go-back/internal/live"
	db "github.com/NeptuneG/go-back/internal/payment/db/sqlc"
	"github.com/NeptuneG/go-back/pkg/db/types"
	"github.com/NeptuneG/go-back/pkg/grpc/interceptors"
	"github.com/NeptuneG/go-back/pkg/log"
	"github.com/dtm-labs/dtm/dtmgrpc"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	dtmGrpcSvrAddr = os.Getenv("DTM_HOST") + ":" + os.Getenv("DTM_GRPC_PORT")
	paymentSvcAddr = os.Getenv("PAYMENT_SERVICE_HOST") + ":" + os.Getenv("PAYMENT_SERVICE_PORT")
	liveSvcAddr    = os.Getenv("LIVE_SERVICE_HOST") + ":" + os.Getenv("LIVE_SERVICE_PORT")
)

type PaymentService struct {
	paymentProto.UnimplementedPaymentServiceServer
	liveClient liveProto.LiveServiceClient
}

func New(ctx context.Context) *PaymentService {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
		grpc.WithUnaryInterceptor(interceptors.ContextPropagatingInterceptor),
	}

	liveClient, err := liveSvc.NewClient(ctx, opts...)
	if err != nil {
		log.Fatal("failed to connect to live service", log.Field.Error(err))
		panic(err)
	}

	return &PaymentService{
		liveClient: liveClient,
	}
}

func (s *PaymentService) Close() {
	if err := db.Close(); err != nil {
		log.Fatal("failed to close database connection", log.Field.Error(err))
		panic(err)
	}
}

func (s *PaymentService) CreateUserPoints(ctx context.Context, req *paymentProto.CreateUserPointsRequest) (*paymentProto.CreateUserPointsResponse, error) {
	log.Info("create user points", log.Field.Any("req", req))

	userID := uuid.MustParse(req.UserId)
	_, err := db.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      userID,
		Points:      req.UserPoints,
		Description: types.NewNullString(req.Description),
		OrderType:   req.OrderType,
		OrderID:     uuid.New(),
		TxID:        types.NewNullString(req.TransactionId),
	})
	if err != nil {
		log.Error("failed to create user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	userPoints, err := db.GetUserPoints(ctx, userID)
	if err != nil {
		log.Error("failed to get user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentProto.CreateUserPointsResponse{
		UserId:     req.UserId,
		UserPoints: userPoints.(int64),
	}, nil
}

func (s *PaymentService) CreateUserPointsCompensate(ctx context.Context, req *paymentProto.CreateUserPointsRequest) (*emptypb.Empty, error) {
	log.Info("create user points compensate", log.Field.Any("req", req))

	if err := db.DeleteUserPointsByTxID(ctx, types.NewNullString(req.TransactionId)); err != nil {
		log.Error("failed to create user points", log.Field.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *PaymentService) GetUserPoints(ctx context.Context, req *paymentProto.GetUserPointsRequest) (*paymentProto.GetUserPointsResponse, error) {
	userID := uuid.MustParse(req.UserId)
	userPoints, err := db.GetUserPoints(ctx, userID)
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
	liveEvent, err := s.validateCreateLiveEventOrderRequest(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	liveEventOrder, err := db.CreateLiveEventOrder(ctx, db.CreateLiveEventOrderParams{
		UserID:      uuid.MustParse(req.UserId),
		LiveEventID: uuid.MustParse(req.LiveEventId),
		Price:       req.Price,
		UserPoints:  req.UserPoints,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	switch {
	case req.Mode == "saga":
		err = s.createLiveEventOrderSaga(req, liveEventOrder.ID, liveEvent.Title)
	case req.Mode == "ttc":
		err = s.createLiveEventOrderTtc(req, liveEvent)
	default:
		err = s.createLiveEventOrderSaga(req, liveEventOrder.ID, liveEvent.Title)
	}

	if err != nil {
		log.Error("failed to create live event order", log.Field.Error(err))
	}

	return &paymentProto.CreateLiveEventOrderResponse{
		State: string(liveEventOrder.State),
	}, nil
}

func (s *PaymentService) SucceedLiveEventOrder(ctx context.Context, req *paymentProto.SucceedLiveEventOrderRequest) (*paymentProto.SucceedLiveEventOrderResponse, error) {
	log.Info("succeed live event order", log.Field.Any("req", req))
	liveEventOrder, err := db.GetLiveEventOrder(ctx, uuid.MustParse(req.LiveEventOrderId))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err := liveEventOrder.UpdateState(ctx, db.StatePaid); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentProto.SucceedLiveEventOrderResponse{
		State: string(liveEventOrder.State),
	}, nil
}

func (s *PaymentService) SucceedLiveEventOrderCompensate(ctx context.Context, req *paymentProto.SucceedLiveEventOrderRequest) (*paymentProto.SucceedLiveEventOrderResponse, error) {
	log.Info("succeed live event order compensate", log.Field.Any("req", req))
	liveEventOrder, err := db.GetLiveEventOrder(ctx, uuid.MustParse(req.LiveEventOrderId))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err := liveEventOrder.UpdateState(ctx, db.StateFailed); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &paymentProto.SucceedLiveEventOrderResponse{
		State: string(liveEventOrder.State),
	}, nil
}

func (s *PaymentService) validateCreateLiveEventOrderRequest(ctx context.Context, req *paymentProto.CreateLiveEventOrderRequest) (*liveProto.LiveEvent, error) {
	userPoints, err := db.GetUserPoints(ctx, uuid.MustParse(req.UserId))
	if userPoints == nil {
		log.Error("user points not found", log.Field.Error(err))
		return nil, errors.New("user points not found")
	}
	if userPoints.(int64) < int64(req.UserPoints) {
		return nil, errors.New("user points not enough")
	}

	liveEventResp, err := s.liveClient.GetLiveEvent(ctx, &liveProto.GetLiveEventRequest{
		Id: req.LiveEventId,
	})
	if liveEventResp.LiveEvent == nil {
		log.Error("failed to get live event", log.Field.Error(err))
		return nil, errors.New("live event not found")
	}
	if liveEventResp.LiveEvent.AvailableSeats == 0 {
		return nil, errors.New("no seats available")
	}
	if liveEventResp.LiveEvent.StageOneStartAt.AsTime().Before(time.Now()) {
		return nil, errors.New("live event has started")
	}
	return liveEventResp.LiveEvent, nil
}

func (s *PaymentService) createLiveEventOrderSaga(req *paymentProto.CreateLiveEventOrderRequest, orderID uuid.UUID, liveEventTitle string) error {
	txGID := dtmgrpc.MustGenGid(dtmGrpcSvrAddr)
	log.Info("init tx", log.Field.String("txGID", txGID))

	saga := dtmgrpc.NewSagaGrpc(dtmGrpcSvrAddr, txGID).
		Add(
			paymentSvcAddr+"/"+paymentProto.PaymentService_ServiceDesc.ServiceName+"/CreateUserPoints",
			paymentSvcAddr+"/"+paymentProto.PaymentService_ServiceDesc.ServiceName+"/CreateUserPointsCompensate",
			&paymentProto.CreateUserPointsRequest{
				UserId:        req.UserId,
				UserPoints:    -req.UserPoints,
				Description:   "Reserve " + liveEventTitle,
				OrderType:     "LiveEventOrder",
				TransactionId: txGID,
			},
		).
		Add(
			liveSvcAddr+"/"+liveProto.LiveService_ServiceDesc.ServiceName+"/ReserveSeat",
			liveSvcAddr+"/"+liveProto.LiveService_ServiceDesc.ServiceName+"/ReserveSeatCompensate",
			&liveProto.ReserveSeatRequest{
				LiveEventId: req.LiveEventId,
			},
		).
		Add(
			paymentSvcAddr+"/"+paymentProto.PaymentService_ServiceDesc.ServiceName+"/SucceedLiveEventOrder",
			paymentSvcAddr+"/"+paymentProto.PaymentService_ServiceDesc.ServiceName+"/SucceedLiveEventOrderCompensate",
			&paymentProto.SucceedLiveEventOrderRequest{
				LiveEventOrderId: orderID.String(),
			},
		).
		EnableConcurrent()

	saga.RetryInterval = 20
	saga.TimeoutToFail = 40

	return saga.Submit()
}

func (s *PaymentService) createLiveEventOrderTtc(req *paymentProto.CreateLiveEventOrderRequest, liveEvent *liveProto.LiveEvent) error {
	return nil
}
