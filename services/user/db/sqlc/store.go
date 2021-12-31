package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NeptuneG/go-back/pkg/types"
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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) (interface{}, error)) (interface{}, error) {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	q := New(tx)
	result, err := fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return nil, err
	}

	return result, tx.Commit()
}

func (store *Store) CreateUserTx(ctx context.Context, param CreateUserParams) (*User, error) {
	if result, err := store.execTx(ctx, func(q *Queries) (interface{}, error) {
		user, err := q.CreateUser(ctx, param)
		if err != nil {
			return nil, err
		}
		_, err = q.CreateUserPoints(ctx, CreateUserPointsParams{
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
