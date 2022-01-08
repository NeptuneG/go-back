// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/NeptuneG/go-back/pkg/db/types"
	"github.com/google/uuid"
)

type LiveEvent struct {
	ID              uuid.UUID        `json:"id"`
	LiveHouseID     uuid.UUID        `json:"live_house_id"`
	Title           string           `json:"title"`
	Url             string           `json:"url"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	Slug            types.NullString `json:"slug"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Seats           int32            `json:"seats"`
	AvailableSeats  int32            `json:"available_seats"`
}

type LiveHouse struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	Address   types.NullString `json:"address"`
	Slug      types.NullString `json:"slug"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
