package main

import (
	"context"
	"net/http"

	proto "github.com/NeptuneG/go-back/api/proto/gateway"
	service "github.com/NeptuneG/go-back/internal/gateway"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"github.com/NeptuneG/go-back/internal/pkg/grpc/interceptors"
	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort = ":3377"
	httpPort = ":4000"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	svc := service.New(ctx)
	grpcSrc := grpcServer.New(grpcPort, func(srv *grpc.Server) {
		proto.RegisterGatewayServiceServer(srv, svc)
	}, grpc.UnaryInterceptor(
		interceptors.UnaryDefaultAuthInterceptor(service.AuthRequiredMethods...),
	))

	go func() {
		grpcSrc.Start()
	}()

	mux := runtime.NewServeMux(
		runtime.WithMetadata(service.PropagateTracingHeader),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := proto.RegisterGatewayServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Error("failed to register gateway", log.Field.Error(err))
		panic(err)
	}

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Error("failed to start server", log.Field.Error(err))
		panic(err)
	}
}
