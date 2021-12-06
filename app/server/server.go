package server

import (
	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/server/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router                  *gin.Engine
	scraped_events_consumer *ScrapedEventsConsumer
}

func NewServer(store *db.Store) *Server {
	controller := controller.NewController(store)
	router := newRouter(controller)
	scraped_consumer := &ScrapedEventsConsumer{store: store}

	return &Server{router: router, scraped_events_consumer: scraped_consumer}
}

func (server *Server) Start(address string) error {
	server.scraped_events_consumer.Start()
	return server.router.Run(address)
}
