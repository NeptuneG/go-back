package db

import (
	"sync"

	"github.com/NeptuneG/go-back/pkg/db"
)

var (
	queriesOnce               sync.Once
	queries                   *Queries
	Close                     = store().Close
	CreateUserPoints          = store().CreateUserPoints
	DeleteUserPointsByTxID    = store().DeleteUserPointsByTxID
	GetUserPoints             = store().GetUserPoints
	CreateLiveEventOrder      = store().CreateLiveEventOrder
	UpdateLiveEventOrderState = store().UpdateLiveEventOrderState
	GetLiveEventOrder         = store().GetLiveEventOrder
)

func store() *Queries {
	queriesOnce.Do(func() {
		queries = New(db.ConnectDatabase())
	})
	return queries
}
