package main

import (
	"database/sql"
	"net"

	"github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	"github.com/NeptuneG/go-back/services/live/server"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/live_development?sslmode=disable"
	port     = ":3377"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close database connection", logField.Error(err))
		}
	}()
	if err != nil {
		log.Fatal("failed to open database connection", logField.Error(err))
	}

	srv := grpc.NewServer()
	proto.RegisterLiveServiceServer(srv, server.New(conn))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", logField.Error(err))
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", logField.Error(err))
	}
}
