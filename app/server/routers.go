package server

import (
	"github.com/NeptuneG/go-back/app/server/controller"
	"github.com/gin-gonic/gin"
)

func newRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/live_houses", controller.GetAllLiveHouses)
		v1.POST("/live_houses", controller.CreateLivehouse)
		v1.POST("/live_events/scrape_jobs", controller.CreateScrapeLiveEventsJob)
		v1.GET("/live_events", controller.GetLiveEvents)
		v1.POST("/users", controller.CreateUser)
		v1.POST("/user_orders", controller.CreateUserOrder)
	}

	return router
}
