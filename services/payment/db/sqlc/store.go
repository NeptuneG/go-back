package db

import (
	"database/sql"

	"github.com/NeptuneG/go-back/pkg/db"
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
