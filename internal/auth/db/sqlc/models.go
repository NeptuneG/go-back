// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/NeptuneG/go-back/internal/pkg/db/types"
	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID        `json:"id"`
	Email              string           `json:"email"`
	EncryptedPassword  string           `json:"encrypted_password"`
	ResetPasswordToken types.NullString `json:"reset_password_token"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}