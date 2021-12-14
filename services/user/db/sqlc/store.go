package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NeptuneG/go-back/pkg/types"
	"github.com/NeptuneG/go-back/services/user/proto"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) CreateUserTx(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	queries := New(tx)
	encrypted_password, err := encryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	arg := CreateUserParams{
		Email:             req.Email,
		EncryptedPassword: encrypted_password,
	}
	user, err := queries.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	user_points, err := queries.CreateUserPoints(ctx, CreateUserPointsParams{
		UserID:      user.ID,
		Points:      1000,
		Description: types.NewNullString("Initial points"),
	})

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return nil, err
	}

	return &proto.CreateUserResponse{
		User: &proto.User{
			Id:     user.ID.String(),
			Email:  user.Email,
			Points: user_points.Points,
		},
	}, tx.Commit()
}

func encryptPassword(password string) (string, error) {
	encrypt_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt_bytes), nil
}
