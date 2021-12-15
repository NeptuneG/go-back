// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/NeptuneG/go-back/pkg/types"
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

type UserOrder struct {
	ID          uuid.UUID `json:"id"`
	LiveEventID uuid.UUID `json:"live_event_id"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserPoint struct {
	ID          uuid.UUID        `json:"id"`
	UserID      uuid.UUID        `json:"user_id"`
	Points      int32            `json:"points"`
	Description types.NullString `json:"description"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}