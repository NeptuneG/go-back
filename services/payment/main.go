package main

import (
	proto "github.com/NeptuneG/go-back/api/proto/payment"
	grpcServer "github.com/NeptuneG/go-back/pkg/grpc"
	"github.com/NeptuneG/go-back/pkg/grpc/interceptors"
	"github.com/NeptuneG/go-back/services/payment/server"
	"google.golang.org/grpc"
)

const (
	port = 3377
)

func main() {
	server := server.New()
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
