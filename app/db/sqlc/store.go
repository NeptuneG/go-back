package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NeptuneG/go-back/pkg/types"
	"github.com/google/uuid"
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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateUserOrderTxParams struct {
	UserID      uuid.UUID `json:"user_id"`
	LiveEventID uuid.UUID `json:"live_event_id"`
	Points      *int32    `json:"points"`
}

func (store *Store) CreateUserOrderTx(ctx context.Context, param CreateUserOrderTxParams) error {
	return store.execTx(ctx, func(q *Queries) error {
		liveEvent, err := q.GetLiveEventById(ctx, param.LiveEventID)
		if err != nil {
			return err
		}
		if liveEvent.AvailableSeats == 0 {
			return errors.New("No available seats")
		}
		if param.Points != nil {
			user_remain_points, err := q.GetUserPoints(ctx, param.UserID)
			if err != nil {
				return err
			}
			if user_remain_points < int64(*param.Points) {
				return errors.New("Not enough points")
			}
			q.CreateUserPoints(ctx, CreateUserPointsParams{
				UserID:      param.UserID,
				Points:      -*param.Points,
				Description: types.NewNullString("Order " + liveEvent.Title),
			})
		}
		_, err = q.CreateUserOrder(ctx, CreateUserOrderParams{
			UserID:      param.UserID,
			LiveEventID: param.LiveEventID,
		})
		if err != nil {
			return err
		}
		q.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
			ID:             param.LiveEventID,
			AvailableSeats: liveEvent.AvailableSeats - 1,
		})

		return nil
	})
}
