package controller

import (
	"net/http"

	db "github.com/NeptuneG/go-back/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserOrderRequest struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	LiveEventID uuid.UUID `json:"live_event_id" binding:"required"`
	Points      *int32    `json:"points"`
}

func (controller *Controller) CreateUserOrder(ctx *gin.Context) {
	var req createUserOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := controller.store.CreateUserOrderTx(ctx, db.CreateUserOrderTxParams{
		UserID:      req.UserID,
		LiveEventID: req.LiveEventID,
		Points:      req.Points,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{"status": "ok"})
}
