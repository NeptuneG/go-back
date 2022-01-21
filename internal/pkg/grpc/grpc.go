package grpc

import (
	"net"
	"net/http"

	"github.com/NeptuneG/go-back/internal/pkg/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func ListenAndServeMetrics(port string) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("metrics server error", log.Field.Error(err))
		panic(err)
	}
}
