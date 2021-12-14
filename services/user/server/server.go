package server

import (
	"context"
	"database/sql"

	db "github.com/NeptuneG/go-back/services/user/db/sqlc"
	"github.com/NeptuneG/go-back/services/user/proto"
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
	result, err := userService.store.CreateUserTx(ctx, db.CreateUserTxParams{
		Email:    req.Email,
		Password: req.Password,
	})
	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:     result.Id,
			Email:  result.Email,
			Points: result.Points,
		},
	}, err
}
