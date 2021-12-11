package service

import (
	"context"
	"errors"

	db "github.com/NeptuneG/go-back/db/sqlc"
)

func CreateUserOrder(ctx context.Context, store *db.Store, param db.CreateUserOrderParams) error {
	return store.ExecTx(ctx, func(q *db.Queries) error {
		availableSeats, err := q.GetLiveEventAvailableSeatsById(ctx, param.LiveEventID)
		if err != nil {
			return err
		}
		if availableSeats == 0 {
			return errors.New("No available seats")
		}
		_, err = q.CreateUserOrder(ctx, param)
		if err != nil {
			return err
		}
		q.UpdateLiveEventAvailableSeatsById(ctx, db.UpdateLiveEventAvailableSeatsByIdParams{
			ID:             param.LiveEventID,
			AvailableSeats: availableSeats - 1,
		})

		return nil
	})
}
