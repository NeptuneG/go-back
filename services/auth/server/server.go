package server

import (
	"context"

	"github.com/NeptuneG/go-back/gen/go/services/auth/proto"
	"github.com/NeptuneG/go-back/pkg/auth"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	db "github.com/NeptuneG/go-back/services/auth/db/sqlc"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	store *db.Store
}

func New() *AuthService {
	return &AuthService{store: db.NewStore()}
}

func (s *AuthService) Close() {
	if err := s.store.Close(); err != nil {
		log.Fatal("failed to close database connection", logField.Error(err))
		panic(err)
	}
}

func (s *AuthService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	encrypted_password, err := encryptPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to encrypt password")
	}
	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Email:             req.Email,
		EncryptedPassword: encrypted_password,
	})
	if err != nil {
		log.Error("failed to create user", logField.Error(err))
		return nil, status.Error(codes.Internal, "failed to create user")
	}
	token, err := auth.CreateToken(user.ID.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create token")
	}
	return &proto.RegisterResponse{
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := s.store.AuthenticateUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	token, err := auth.CreateToken(user.ID.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create token")
	}
	return &proto.LoginResponse{
		Token: token,
	}, nil
}

func encryptPassword(password string) (string, error) {
	encrypt_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt_bytes), nil
}
