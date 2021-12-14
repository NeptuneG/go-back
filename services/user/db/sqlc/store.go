package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NeptuneG/go-back/pkg/types"
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

type CreateUserTxParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserTxResult struct {
	Id     string
	Email  string
	Points int32
}

var nullCreateUserTxResult CreateUserTxResult

func (store *Store) CreateUserTx(ctx context.Context, param CreateUserTxParams) (CreateUserTxResult, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nullCreateUserTxResult, err
	}

	queries := New(tx)
	encrypted_password, err := encryptPassword(param.Password)
	if err != nil {
		return nullCreateUserTxResult, err
	}
	arg := CreateUserParams{
		Email:             param.Email,
		EncryptedPassword: encrypted_password,
	}
	user, err := queries.CreateUser(ctx, arg)
	if err != nil {
		return nullCreateUserTxResult, err
	}
	user_points, err := queries.CreateUserPoints(ctx, CreateUserPointsParams{
		UserID:      user.ID,
		Points:      1000,
		Description: types.NewNullString("Initial points"),
	})

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nullCreateUserTxResult, fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return nullCreateUserTxResult, err
	}

	return CreateUserTxResult{
		Id:     user.ID.String(),
		Email:  user.Email,
		Points: user_points.Points,
	}, tx.Commit()
}

func encryptPassword(password string) (string, error) {
	encrypt_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt_bytes), nil
}
