package main

import (
	"github.com/NeptuneG/go-back/gen/go/services/scraper/proto"
	grpcServer "github.com/NeptuneG/go-back/pkg/grpc"
	"github.com/NeptuneG/go-back/services/scraper/server"
	"google.golang.org/grpc"
)

const (
	port = 3377
)

func main() {
	server := server.New()
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterScrapeServiceServer(srv, server)
	})

	gprcSrv.Start()
}
