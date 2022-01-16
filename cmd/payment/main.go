package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/payment"
	"github.com/NeptuneG/go-back/internal/payment"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":3377"
)

func main() {
	server := payment.New()
	defer server.Close()

	gprcSrv := grpcServer.New(port, func(srv *grpc.Server) {
		proto.RegisterPaymentServiceServer(srv, server)
	})

	gprcSrv.Start()
}
