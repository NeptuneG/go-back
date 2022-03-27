package db

import (
	"sync"

	"github.com/NeptuneG/go-back/pkg/db"
)

var (
	queriesOnce sync.Once
	queries     *Queries
	Close       = store().Close
)

func store() *Queries {
	queriesOnce.Do(func() {
		queries = New(db.ConnectDatabase())
	})
	return queries
}
