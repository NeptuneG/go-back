package grpc

import (
	"net"

	"github.com/NeptuneG/go-back/internal/pkg/log"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port   string
}

func New(port string, register func(server *grpc.Server), opt ...grpc.ServerOption) *Server {
	srv := grpc.NewServer(opt...)
	register(srv)
	return &Server{
		server: srv,
		port:   port,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatal("failed to listen", log.Field.Error(err))
		panic(err)
	}

	err = s.server.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", log.Field.Error(err))
		panic(err)
	}
}
