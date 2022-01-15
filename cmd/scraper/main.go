package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/scraper"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"github.com/NeptuneG/go-back/internal/scraper"
	"google.golang.org/grpc"
)

const (
	port = ":3377"
)

func main() {
	server := scraper.New()
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterScrapeServiceServer(srv, server)
	})

	gprcSrv.Start()
}
