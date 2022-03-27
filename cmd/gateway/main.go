package main

import (
	"context"
	"net/http"

	proto "github.com/NeptuneG/go-back/api/proto/gateway"
	service "github.com/NeptuneG/go-back/internal/gateway"
	grpcServer "github.com/NeptuneG/go-back/pkg/grpc"

	"github.com/NeptuneG/go-back/pkg/grpc/interceptors"
	"github.com/NeptuneG/go-back/pkg/log"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	grpcPort    = ":3377"
	httpPort    = ":4000"
	metricsPort = ":9887"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	svc := service.New(ctx)
	grpcSrc := grpcServer.New(grpcPort, func(srv *grpc.Server) {
		proto.RegisterGatewayServiceServer(srv, svc)
	}, grpc_middleware.WithUnaryServerChain(
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
		interceptors.UnaryDefaultAuthInterceptor(service.AuthRequiredMethods...),
	))

	go func() { grpcSrc.Start() }()
	go func() { grpcServer.ListenAndServeMetrics(metricsPort) }()

	mux := runtime.NewServeMux(
		runtime.WithMetadata(service.PropagateTracingHeader),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := proto.RegisterGatewayServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Error("failed to register gateway", log.Field.Error(err))
		panic(err)
	}

	if err := http.ListenAndServe(httpPort, mux); err != nil {
		log.Error("failed to start server", log.Field.Error(err))
		panic(err)
	}
}
