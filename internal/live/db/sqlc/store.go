package db

import (
	"context"
	"database/sql"
	"errors"
	"sync"
	"time"

	"github.com/NeptuneG/go-back/internal/pkg/cache"
	"github.com/NeptuneG/go-back/internal/pkg/db"
	"github.com/NeptuneG/go-back/internal/pkg/db/types"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/google/uuid"
)

var (
	queriesOnce             sync.Once
	queries                 *Queries
	Close                   = store().Close
	CreateLiveHouse         = store().CreateLiveHouse
	CreateLiveEvent         = store().CreateLiveEvent
	GetAllLiveHousesDefault = store().GetAllLiveHousesDefault
	GetAllLiveEvents        = store().GetAllLiveEvents
)

func store() *Queries {
	queriesOnce.Do(func() {
		queries = New(db.ConnectDatabase())
	})
	return queries
}

func DecrementLiveEventAvailableSeats(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if result, err := db.ExecTx(ctx, func(tx *sql.Tx) (interface{}, error) {
		q := store().WithTx(tx)
		liveEvent, err := q.GetLiveEventById(ctx, liveEventID)
		if err != nil {
			return nil, err
		}
		if liveEvent.AvailableSeats == 0 {
			return nil, errors.New("no available seats")
		}
		liveEvent.AvailableSeats--
		if err = q.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
			ID:             liveEventID,
			AvailableSeats: liveEvent.AvailableSeats,
		}); err != nil {
			return nil, err
		}
		if err := cacheLiveEvent(ctx, liveEvent); err != nil {
			log.Error("failed to update live event cache", log.Field.Error(err))
		}
		return &liveEvent, nil
	}); err != nil {
		return nil, err
	} else {
		return result.(*GetLiveEventByIdRow), nil
	}
}

func IncrementLiveEventAvailableSeats(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if result, err := db.ExecTx(ctx, func(tx *sql.Tx) (interface{}, error) {
		q := store().WithTx(tx)
		liveEvent, err := q.GetLiveEventById(ctx, liveEventID)
		if err != nil {
			return nil, err
		}
		if liveEvent.AvailableSeats == liveEvent.Seats {
			return &liveEvent, nil
		}
		liveEvent.AvailableSeats++
		if err = q.UpdateLiveEventAvailableSeatsById(ctx, UpdateLiveEventAvailableSeatsByIdParams{
			ID:             liveEventID,
			AvailableSeats: liveEvent.AvailableSeats,
		}); err != nil {
			return nil, err
		}
		if err := cacheLiveEvent(ctx, liveEvent); err != nil {
			log.Error("failed to update live event cache", log.Field.Error(err))
		}
		return &liveEvent, err
	}); err != nil {
		return nil, err
	} else {
		return result.(*GetLiveEventByIdRow), nil
	}
}

func GetLiveHouseBySlug(ctx context.Context, liveHouseSlug string) (*LiveHouse, error) {
	if liveHouseSlug == "" {
		return nil, errors.New("liveHouseSlug is empty")
	}

	var liveHouse LiveHouse
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   "live-house:" + liveHouseSlug,
		Value: &liveHouse,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			return store().GetLiveHouseBySlug(ctx, types.NewNullString(liveHouseSlug))
		},
	}); err != nil {
		return nil, err
	}
	return &liveHouse, nil
}

func GetLiveEventById(ctx context.Context, liveEventID uuid.UUID) (*GetLiveEventByIdRow, error) {
	if liveEventID == uuid.Nil {
		return nil, errors.New("liveEventID is empty")
	}

	var liveEvent GetLiveEventByIdRow
	if err := cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   "live-event:" + liveEventID.String(),
		Value: &liveEvent,
		TTL:   5 * time.Minute,
		Do: func(item *cache.Item) (interface{}, error) {
			liveEvent, err := store().GetLiveEventById(ctx, liveEventID)
			if err != nil {
				return nil, err
			}
			return liveEvent, nil
		},
	}); err != nil {
		return nil, err
	}
	return &liveEvent, nil
}

func cacheLiveEvent(ctx context.Context, liveEvent GetLiveEventByIdRow) error {
	key := "live-event:" + liveEvent.ID.String()
	if err := cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: liveEvent,
	}); err != nil {
		return cache.Delete(ctx, key)
	}
	return nil
}
