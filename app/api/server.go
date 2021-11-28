package api

import (
	"database/sql"
	"net/http"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/live_houses", server.getAllLiveHouses)
	router.POST("/live_houses", server.createLivehouse)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type getAllLiveHousesRequest struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (server *Server) getAllLiveHouses(ctx *gin.Context) {
	var req getAllLiveHousesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllLiveHousesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	liveHouses, err := server.store.GetAllLiveHouses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, liveHouses)
}

type createLiveHouseRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Slug    string `json:"slug"`
}

func (server *Server) createLivehouse(ctx *gin.Context) {
	var req createLiveHouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLiveHouseParams{
		Name:    req.Name,
		Address: sql.NullString{Valid: true, String: req.Address},
		Slug:    sql.NullString{Valid: true, String: req.Slug},
	}
	liveHouse, err := server.store.CreateLiveHouse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, liveHouse)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
