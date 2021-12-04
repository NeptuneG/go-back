package server

import (
	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/server/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	controller := controller.NewController(store)
	router := newRouter(controller)

	return &Server{router: router}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
