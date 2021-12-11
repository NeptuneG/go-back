package controller

import (
	"net/http"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/NeptuneG/go-back/server/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserOrderRequest struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	LiveEventID uuid.UUID `json:"live_event_id" binding:"required"`
}

func (controller *Controller) CreateUserOrder(ctx *gin.Context) {
	var req createUserOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := service.CreateUserOrder(ctx, controller.store, db.CreateUserOrderParams{
		UserID:      req.UserID,
		LiveEventID: req.LiveEventID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{"status": "ok"})
}
