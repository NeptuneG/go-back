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
	return userService.store.CreateUserTx(ctx, req)
}