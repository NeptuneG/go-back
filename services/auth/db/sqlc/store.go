package db

import (
	"context"
	"database/sql"

	"github.com/NeptuneG/go-back/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore() *Store {
	dbConn := db.ConnectDatabase()
	return &Store{
		db:      dbConn,
		Queries: New(dbConn),
	}
}

func (s *Store) AuthenticateUser(ctx context.Context, email string, password string) (*User, error) {
	user, err := s.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)); err != nil {
		return nil, err
	}
	return &user, nil
}
