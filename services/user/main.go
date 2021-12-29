package main

import (
	"database/sql"
	"net"

	"github.com/NeptuneG/go-back/gen/go/services/user/proto"
	"github.com/NeptuneG/go-back/pkg/logger"
	"github.com/NeptuneG/go-back/services/user/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://dev@db/user_development?sslmode=disable"
	port     = ":3377"
)

func main() {
	logger := logger.New()
	grpc_zap.ReplaceGrpcLogger(logger)

	conn, err := sql.Open(dbDriver, dbSource)
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Fatal("failed to close database connection", zap.Error(err))
		}
	}()
	if err != nil {
		logger.Fatal("failed to open database connection", zap.Error(err))
		return
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)
	proto.RegisterUserServiceServer(srv, server.New(conn, logger))

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
