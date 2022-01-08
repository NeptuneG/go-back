package db

import (
	"context"
	"database/sql"

	"github.com/NeptuneG/go-back/pkg/db"
	"github.com/NeptuneG/go-back/pkg/db/types"
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

func (store *Store) CreateUserTx(ctx context.Context, param CreateUserParams) (*User, error) {
	if result, err := db.ExecTx(ctx, func(tx *sql.Tx) (interface{}, error) {
		queries := store.Queries.WithTx(tx)
		user, err := queries.CreateUser(ctx, param)
		if err != nil {
			return nil, err
		}
		_, err = queries.CreateUserPoints(ctx, CreateUserPointsParams{
			UserID:      user.ID,
			Points:      1000,
			Description: types.NewNullString("Initial points"),
			OrderID:     types.NewNullUUID(nil),
		})
		return &user, err
	}); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}
