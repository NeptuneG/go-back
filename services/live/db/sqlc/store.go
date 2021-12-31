package db

import (
	"context"
	"database/sql"
	"fmt"

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

func (store *Store) ReserveSeatTx(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if result, err := store.execTx(ctx, func(q *Queries) (interface{}, error) {
		liveEvent, err := q.GetLiveEventById(ctx, liveEventID)
		if err != nil {
			return nil, err
		}
		if liveEvent.AvailableSeats == 0 {
			return nil, fmt.Errorf("no available seats")
		}
		liveEvent.AvailableSeats--
		if err = q.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
			ID:             liveEventID,
			AvailableSeats: liveEvent.AvailableSeats,
		}); err != nil {
			return nil, err
		}
		return &liveEvent, nil
	}); err != nil {
		return nil, err
	} else {
		return result.(*GetLiveEventByIdRow), nil
	}
}

func (store *Store) RollbackSeatReservationTx(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if result, err := store.execTx(ctx, func(q *Queries) (interface{}, error) {
		liveEvent, err := q.GetLiveEventById(ctx, liveEventID)
		if err != nil {
			return nil, err
		}
		liveEvent.AvailableSeats++
		if err = q.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
			ID:             liveEventID,
			AvailableSeats: liveEvent.AvailableSeats,
		}); err != nil {
			return nil, err
		}
		return &liveEvent, err
	}); err != nil {
		return nil, err
	} else {
		return result.(*GetLiveEventByIdRow), nil
	}
}
