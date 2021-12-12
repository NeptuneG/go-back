package controller

import (
	"net/http"

	db "github.com/NeptuneG/go-back/app/db/sqlc"
	"github.com/NeptuneG/go-back/pkg/types"
	"github.com/gin-gonic/gin"
)

type getAllLiveHousesRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (controller *Controller) GetAllLiveHouses(ctx *gin.Context) {
	var req getAllLiveHousesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	liveHouses, err := controller.store.GetAllLiveHousesDefault(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string][]db.GetAllLiveHousesDefaultRow{"live_houses": liveHouses})
}

type createLiveHouseRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Slug    string `json:"slug"`
}

func (controller *Controller) CreateLivehouse(ctx *gin.Context) {
	var req createLiveHouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLiveHouseParams{
		Name:    req.Name,
		Address: types.NewNullString(req.Address),
		Slug:    types.NewNullString(req.Slug),
	}
	liveHouse, err := controller.store.CreateLiveHouse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, liveHouse)
}
