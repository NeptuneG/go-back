package main

import (
	"context"
	"database/sql"
	"net"

	"github.com/NeptuneG/go-back/gen/go/services/payment/proto"
	"github.com/NeptuneG/go-back/pkg/logger"
	"github.com/NeptuneG/go-back/services/payment/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/payment_development?sslmode=disable"
	port     = ":3377"
)

func main() {
	logger := logger.New()

	dbConn, err := sql.Open(dbDriver, dbSource)
	defer func() {
		if err := dbConn.Close(); err != nil {
			logger.Fatal("failed to close database connection", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("failed to open database connection", zap.Error(err))
		return
	}

	srv := grpc.NewServer()
	ctx := context.Background()
	service, err := server.New(ctx, dbConn, logger)
	if err != nil {
		logger.Fatal("failed to create payment service", zap.Error(err))
		return
	}
	proto.RegisterPaymentServiceServer(srv, service)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
		return
	}

	err = srv.Serve(listener)
	if err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
		return
	}
}
