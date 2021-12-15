package main

import (
	"context"
	"log"
	"net"

	live "github.com/NeptuneG/go-back/services/live/proto"
	"github.com/NeptuneG/go-back/services/scraper/consumer"
	"github.com/NeptuneG/go-back/services/scraper/proto"
	"github.com/NeptuneG/go-back/services/scraper/server"
	"google.golang.org/grpc"
)

const (
	port = ":3377"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "live-service:3377", opts...)
	if err != nil {
		log.Fatalf("failed to dial grpc server: %v", err)
		return
	}

	liveClient := live.NewLiveServiceClient(conn)

	consumer := consumer.New(liveClient)
	consumer.Start(ctx)

	srv := grpc.NewServer()
	proto.RegisterScrapeServiceServer(srv, &server.ScrapeService{})

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
