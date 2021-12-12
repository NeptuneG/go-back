package server

import (
	db "github.com/NeptuneG/go-back/app/db/sqlc"
	"github.com/NeptuneG/go-back/app/server/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router                *gin.Engine
	scrapedEventsConsumer *ScrapedEventsConsumer
}

func NewServer(store *db.Store) *Server {
	controller := controller.NewController(store)
	router := newRouter(controller)
	consumer := &ScrapedEventsConsumer{store: store}

	return &Server{router: router, scrapedEventsConsumer: consumer}
}

func (server *Server) Start(address string) error {
	server.scrapedEventsConsumer.Start()
	return server.router.Run(address)
}
