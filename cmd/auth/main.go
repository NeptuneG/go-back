package main

import (
	"context"

	proto "github.com/NeptuneG/go-back/api/proto/auth"
	"github.com/NeptuneG/go-back/internal/auth"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":3377"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := auth.New(ctx)
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterAuthServiceServer(srv, server)
	})

	gprcSrv.Start()
}
