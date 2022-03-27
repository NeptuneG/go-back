package auth

import (
	"context"

	proto "github.com/NeptuneG/go-back/api/proto/auth"
	db "github.com/NeptuneG/go-back/internal/auth/db/sqlc"
	"github.com/NeptuneG/go-back/pkg/auth"
	"github.com/NeptuneG/go-back/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
}

func New(ctx context.Context) *AuthService {
	return &AuthService{}
}

func (s *AuthService) Close() {
	if err := db.Close(); err != nil {
		log.Fatal("failed to close database connection", log.Field.Error(err))
		panic(err)
	}
}

func (s *AuthService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	user, err := db.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Error("failed to create user", log.Field.Error(err))
		return nil, status.Error(codes.Internal, "failed to create user")
	}
	token, err := auth.CreateToken(user.ID.String())
	if err != nil {
		log.Error("failed to create toekn for user", log.Field.Error(err))
		return nil, status.Error(codes.Internal, "failed to create token")
	}
	return &proto.RegisterResponse{
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := db.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if err = user.Authenticate(req.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "password is incorrect")
	}
	token, err := auth.CreateToken(user.ID.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create token")
	}
	return &proto.LoginResponse{
		Token: token,
	}, nil
}
