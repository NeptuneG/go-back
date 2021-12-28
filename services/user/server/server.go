package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/user/db/sqlc"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	count = 0
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	store  *db.Store
	logger *zap.Logger
}

func New(dbConn *sql.DB, logger *zap.Logger) *UserService {
	return &UserService{
		store:  db.NewStore(dbConn),
		logger: logger,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.store.CreateUserTx(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	user, err := s.store.GetUserByID(ctx, userID)
	return &proto.GetUserResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user.Points,
		},
	}, err
}

func (s *UserService) ConsumeUserPoints(ctx context.Context, req *proto.ConsumeUserPointsRequest) (*proto.ConsumeUserPointsResponse, error) {
	// force a retry
	count++
	s.logger.Debug("mock failure for retry", zap.Int("count", count))
	if count%3 != 0 {
		return nil, status.Error(codes.Internal, "just failed")
	}

	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	_, err = s.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      userID,
		Points:      -req.Points,
		Description: types.NewNullString(req.Description),
		OrderID:     types.NewNullUUID(&orderID),
	})
	if err != nil {
		return nil, err
	}

	// mock delay
	if false {
		s.logger.Debug("mock delay")
		time.Sleep(10 * time.Second)
		s.logger.Debug("mock delay done")
	}

	user, err := s.store.GetUserByID(ctx, userID)
	return &proto.ConsumeUserPointsResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user.Points,
		},
	}, err
}

func (s *UserService) RollbackConsumeUserPoints(ctx context.Context, req *proto.RollbackConsumeUserPointsRequest) (*proto.RollbackConsumeUserPointsResponse, error) {
	orderId, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}

	if err := s.store.DeleteUserPointsByOrderID(ctx, types.NewNullUUID(&orderId)); err != nil {
		return nil, err
	}

	return &proto.RollbackConsumeUserPointsResponse{
		Success: true,
	}, nil
}
