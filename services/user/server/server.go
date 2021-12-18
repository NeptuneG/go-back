package server

import (
	"context"
	"database/sql"

	"github.com/NeptuneG/go-back/gen/go/services/user/proto"
	db "github.com/NeptuneG/go-back/services/user/db/sqlc"
	"github.com/google/uuid"
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
			Points: int32(user.Points),
		},
	}, err
}

func (userService *UserService) IsUserExist(ctx context.Context, req *proto.IsUserExistRequest) (*proto.IsUserExistResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	exist, err := userService.store.IsUserExist(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &proto.IsUserExistResponse{
		Exist: exist,
	}, nil
}
