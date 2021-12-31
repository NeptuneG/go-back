package grpc

import (
	"fmt"
	"net"

	"github.com/NeptuneG/go-back/pkg/log"
	logField "github.com/NeptuneG/go-back/pkg/log/field"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port   int
}

func New(port int, register func(server *grpc.Server)) *Server {
	srv := grpc.NewServer()
	register(srv)
	return &Server{
		server: srv,
		port:   port,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatal("failed to listen", logField.Error(err))
		panic(err)
	}

	err = s.server.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", logField.Error(err))
		panic(err)
	}
}
