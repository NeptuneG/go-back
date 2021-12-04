package server

import (
	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/server/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	pubsub *PubSub
}

func NewServer(store *db.Store) *Server {
	controller := controller.NewController(store)
	router := newRouter(controller)
	pubsub := &PubSub{store: store}

	return &Server{router: router, pubsub: pubsub}
}

func (server *Server) Start(address string) error {
	server.pubsub.StartPubSub()
	return server.router.Run(address)
}
