package db

import (
	"context"
	"sync"

	"github.com/NeptuneG/go-back/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	queriesOnce    sync.Once
	queries        *Queries
	Close          = store().Close
	GetUserByEmail = store().GetUserByEmail
)

func store() *Queries {
	queriesOnce.Do(func() {
		queries = New(db.ConnectDatabase())
	})
	return queries
}

func CreateUser(ctx context.Context, email string, password string) (*User, error) {
	encrypted_password, err := encryptPassword(password)
	if err != nil {
		return nil, err
	}

	if user, err := store().CreateUser(ctx, CreateUserParams{
		Email:             email,
		EncryptedPassword: encrypted_password,
	}); err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func encryptPassword(password string) (string, error) {
	encrypt_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt_bytes), nil
}
