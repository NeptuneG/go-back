package main

import (
	"context"

	proto "github.com/NeptuneG/go-back/api/proto/live"
	"github.com/NeptuneG/go-back/internal/live"
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
	server := live.New(context.Background())
	defer server.Close()

	go func() { grpcServer.ListenAndServeMetrics(metricsPort) }()

	gprcSrv := grpcServer.New(grpcPort, func(srv *grpc.Server) {
		proto.RegisterLiveServiceServer(srv, server)
	}, grpc_middleware.WithUnaryServerChain(
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	))

	gprcSrv.Start()
}
