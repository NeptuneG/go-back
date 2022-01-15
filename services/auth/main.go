package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/auth"
	grpcServer "github.com/NeptuneG/go-back/pkg/grpc"
	"github.com/NeptuneG/go-back/services/auth/server"
	"google.golang.org/grpc"
)

const (
	port = 3377
)

func main() {
	server := server.New()
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterAuthServiceServer(srv, server)
	})

	gprcSrv.Start()
}
