package controller

import (
	"net/http"
	"strconv"

	db "github.com/NeptuneG/go-back/app/db/sqlc"
	"github.com/NeptuneG/go-back/pkg/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *Controller) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	encrypted_password, err := encryptPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:             req.Email,
		EncryptedPassword: encrypted_password,
	}
	user, err := controller.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	user_points, err := controller.store.CreateUserPoints(ctx, db.CreateUserPointsParams{
		UserID:      user.ID,
		Points:      1000,
		Description: types.NewNullString("Initial points"),
	})

	ctx.JSON(http.StatusCreated, map[string]string{
		"id":     user.ID.String(),
		"email":  user.Email,
		"points": strconv.Itoa(int(user_points.Points)),
	})
}

func encryptPassword(password string) (string, error) {
	encrypt_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt_bytes), nil
}
