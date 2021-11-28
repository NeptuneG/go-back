package main

import (
	"database/sql"
	"log"

	"github.com/NeptuneG/go-back/api"
	db "github.com/NeptuneG/go-back/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://dev@db/back_development?sslmode=disable"
	serverAddress = "0.0.0.0:3000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Connecting database failed:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Starting server failed:", err)
	}
}
