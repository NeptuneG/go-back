package db

import (
	"context"

	"github.com/NeptuneG/go-back/pkg/log"
)

func (lo *LiveEventOrder) UpdateState(ctx context.Context, state State) error {
	err := UpdateLiveEventOrderState(ctx, UpdateLiveEventOrderStateParams{
		ID:    lo.ID,
		State: state,
	})
	if err != nil {
		log.Error("update live event order state failed", log.Field.Error(err))
	} else {
		lo.State = state
	}
	return err
}

func (lo *LiveEventOrder) Reload(ctx context.Context) *LiveEventOrder {
	liveOrder, err := GetLiveEventOrder(ctx, lo.ID)
	if err != nil {
		log.Error("failed to reload live event order", log.Field.Error(err))
	} else {
		lo = &liveOrder
	}
	return lo
}
