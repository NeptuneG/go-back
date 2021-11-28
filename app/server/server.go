package server

import (
	"github.com/NeptuneG/go-back/controller"
	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	controller *controller.Controller
}

func NewServer(store *db.Store) *Server {
	router := gin.Default()
	controller := controller.NewController(store)

	router.GET("/live_houses", controller.GetAllLiveHouses)
	router.POST("/live_houses", controller.CreateLivehouse)

	return &Server{
		router:     router,
		controller: controller,
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
