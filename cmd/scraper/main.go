package main

import (
	"context"

	proto "github.com/NeptuneG/go-back/api/proto/scraper"
	"github.com/NeptuneG/go-back/internal/scraper"
	"github.com/NeptuneG/go-back/internal/scraper/consumer"
	grpcServer "github.com/NeptuneG/go-back/pkg/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

const (
	grpcPort    = ":3377"
	metricsPort = ":9887"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() { grpcServer.ListenAndServeMetrics(metricsPort) }()
	go func() { consumer.New(ctx).Start() }()

	server := scraper.New(ctx)
	defer server.Close()

	gprcSrv := grpcServer.New(grpcPort, func(srv *grpc.Server) {
		proto.RegisterScrapeServiceServer(srv, server)
	}, grpc_middleware.WithUnaryServerChain(
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	))

	gprcSrv.Start()
}
