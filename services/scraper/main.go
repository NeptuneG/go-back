package main

import (
	"context"
	"net"

	live "github.com/NeptuneG/go-back/gen/go/services/live/proto"
	"github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	"github.com/NeptuneG/go-back/pkg/logger"
	"github.com/NeptuneG/go-back/services/scraper/consumer"
	"github.com/NeptuneG/go-back/services/scraper/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = ":3377"
)

func main() {
	logger := logger.New()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		logger.Fatal("failed to dial live-service", zap.Error(err))
		return
	}

	liveClient := live.NewLiveServiceClient(conn)

	consumer := consumer.New(liveClient, logger)
	consumer.Start(ctx)

	srv := grpc.NewServer()
	proto.RegisterScrapeServiceServer(srv, &server.ScrapeService{})

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
