package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/services/live/server"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@live-db/live_development?sslmode=disable"
	port     = ":3377"
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

	srv := grpc.NewServer()
	proto.RegisterLiveServiceServer(srv, server.New(conn))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
