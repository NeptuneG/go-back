package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/payment"
	"github.com/NeptuneG/go-back/internal/payment"
	grpcServer "github.com/NeptuneG/go-back/internal/pkg/grpc"
	"github.com/NeptuneG/go-back/internal/pkg/grpc/interceptors"
	"google.golang.org/grpc"
)

const (
	port = 3377
)

func main() {
	server := payment.New()
	defer server.Close()

	register := func(srv *grpc.Server) {
		proto.RegisterPaymentServiceServer(srv, server)
	}
	authInterceptor := interceptors.UnaryDefaultAuthInterceptor("*")
	gprcSrv := grpcServer.New(
		port,
		register,
		grpc.UnaryInterceptor(authInterceptor),
	)

	gprcSrv.Start()
}
