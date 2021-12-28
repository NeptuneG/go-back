package server

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/types"
	db "github.com/NeptuneG/go-back/services/user/db/sqlc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	count = 0
)

type UserService struct {
	proto.UnimplementedUserServiceServer
	store *db.Store
}

func New(dbConn *sql.DB) *UserService {
	return &UserService{
		store: db.NewStore(dbConn),
	}
}

func (userService *UserService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	return userService.store.CreateUserTx(ctx, req)
}

func (userService *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	user, err := userService.store.GetUserByID(ctx, userID)
	return &proto.GetUserResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user.Points,
		},
	}, err
}

func (userService *UserService) ConsumeUserPoints(ctx context.Context, req *proto.ConsumeUserPointsRequest) (*proto.ConsumeUserPointsResponse, error) {
	// force a retry
	count++
	log.Print("count: ", count)
	if count%3 != 0 {
		return nil, status.Error(codes.Internal, "just failed")
	}

	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = userService.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      userID,
		Points:      -req.Points,
		Description: types.NewNullString(req.Description),
	})
	if err != nil {
		return nil, err
	}

	// mock delay
	if false {
		log.Println("mock delay")
		time.Sleep(10 * time.Second)
		log.Println("mock delay done")
	}

	user, err := userService.store.GetUserByID(ctx, userID)
	return &proto.ConsumeUserPointsResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user.Points,
		},
	}, err
}
