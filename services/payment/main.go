package main

import (
	"context"
	"database/sql"
	"net"

	"github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	"github.com/NeptuneG/go-back/services/payment/server"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/payment_development?sslmode=disable"
	port     = ":3377"
)

func main() {
	dbConn, err := sql.Open(dbDriver, dbSource)
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Fatal("failed to close database connection", logField.Error(err))
		}
	}()
	if err != nil {
		log.Fatal("failed to open database connection", logField.Error(err))
		return
	}

	srv := grpc.NewServer()
	ctx := context.Background()
	service, err := server.New(ctx, dbConn)
	if err != nil {
		log.Fatal("failed to create payment service", logField.Error(err))
		return
	}
	proto.RegisterPaymentServiceServer(srv, service)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", logField.Error(err))
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", logField.Error(err))
		return
	}
}
