package main

import (
	"database/sql"
	"log"

	db "github.com/NeptuneG/go-back/app/db/sqlc"
	"github.com/NeptuneG/go-back/app/server"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://dev@db/back_development?sslmode=disable"
	serverAddress = "0.0.0.0:3000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal("Connecting database failed:", err)
	}

	store := db.NewStore(conn)
	server := server.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Starting server failed:", err)
	}
}
