package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NeptuneG/go-back/internal/pkg/cache"
	"github.com/NeptuneG/go-back/internal/pkg/db"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/google/uuid"
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

func (store *Store) ReserveSeatTx(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if result, err := db.ExecTx(ctx, func(tx *sql.Tx) (interface{}, error) {
		q := store.Queries.WithTx(tx)
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
		if err := store.updateLiveEventCache(ctx, liveEvent); err != nil {
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
	if result, err := db.ExecTx(ctx, func(tx *sql.Tx) (interface{}, error) {
		q := store.Queries.WithTx(tx)
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
		if err := store.updateLiveEventCache(ctx, liveEvent); err != nil {
			log.Error("failed to update live event cache", log.Field.Error(err))
		}
		return &liveEvent, err
	}); err != nil {
		return nil, err
	} else {
		return result.(*GetLiveEventByIdRow), nil
	}
}

func (store *Store) updateLiveEventCache(ctx context.Context, liveEvent GetLiveEventByIdRow) error {
	if err := cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf("live-event:%s", liveEvent.ID.String()),
		Value: liveEvent,
	}); err != nil {
		if err := cache.Delete(ctx, fmt.Sprintf("live-event:%s", liveEvent.ID.String())); err != nil {
			return err
		}
		return err
	}
	return nil
}
