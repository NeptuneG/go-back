package main

import (
	"context"

	proto "github.com/NeptuneG/go-back/api/proto/live"
	"github.com/NeptuneG/go-back/internal/live"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
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
	})

	gprcSrv.Start()
}
