package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/live"
	"github.com/NeptuneG/go-back/internal/live"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	port = 3377
)

func main() {
	server := live.New()
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterLiveServiceServer(srv, server)
	})

	gprcSrv.Start()
}
