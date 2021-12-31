package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/user/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/user_development?sslmode=disable"
)

var (
	count = 0
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	store  *db.Store
	dbConn *sql.DB
}

func New() *UserService {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("failed to open database connection", logField.Error(err))
		panic(err)
	}
	return &UserService{store: db.NewStore(dbConn), dbConn: dbConn}
}

func (s *UserService) Close() {
	if err := s.store.Close(); err != nil {
		log.Fatal("failed to close database connection", logField.Error(err))
		panic(err)
	}
	if err := s.dbConn.Close(); err != nil {
		log.Fatal("failed to close database connection", logField.Error(err))
		panic(err)
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return s.store.CreateUserTx(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse id")
	}
	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	return &proto.GetUserResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user.Points,
		},
	}, nil
}

func (s *UserService) ConsumeUserPoints(ctx context.Context, req *proto.ConsumeUserPointsRequest) (*proto.ConsumeUserPointsResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse user id")
	}
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse order id")
	}

	// force a retry
	count++
	log.Debug("mock failure for retry", logField.Int("count", count))
	if count%3 != 0 {
		return nil, status.Error(codes.Internal, "just failed")
	}

	_, err = s.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      userID,
		Points:      -req.Points,
		Description: types.NewNullString(req.Description),
		OrderID:     types.NewNullUUID(&orderID),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to consume user points")
	}

	// mock delay
	if false {
		log.Debug("mock delay")
		time.Sleep(10 * time.Second)
		log.Debug("mock delay done")
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
		return nil, status.Error(codes.InvalidArgument, "failed to parse order id")
	}

	if err := s.store.DeleteUserPointsByOrderID(ctx, types.NewNullUUID(&orderId)); err != nil {
		return nil, status.Error(codes.Internal, "failed to rollback user points")
	}

	return &proto.RollbackConsumeUserPointsResponse{
		Success: true,
	}, nil
}
