// Code generated by sqlc. DO NOT EDIT.
// source: live_events.sql

package db

import (
	"context"
	"time"

	"github.com/NeptuneG/go-back/pkg/db/types"
	"github.com/google/uuid"
)

const createLiveEvent = `-- name: CreateLiveEvent :one
INSERT INTO live_events (
  live_house_id, title, url,
  description, price_info,
  stage_one_open_at, stage_one_start_at,
  stage_two_open_at, stage_two_start_at,
  seats, available_seats
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING id, live_house_id, title, url, description, price_info, stage_one_open_at, stage_one_start_at, stage_two_open_at, stage_two_start_at, slug, created_at, updated_at, seats, available_seats
`

type CreateLiveEventParams struct {
	LiveHouseID     uuid.UUID        `json:"live_house_id"`
	Title           string           `json:"title"`
	Url             string           `json:"url"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	Seats           int32            `json:"seats"`
	AvailableSeats  int32            `json:"available_seats"`
}

func (q *Queries) CreateLiveEvent(ctx context.Context, arg CreateLiveEventParams) (LiveEvent, error) {
	row := q.queryRow(ctx, q.createLiveEventStmt, createLiveEvent,
		arg.LiveHouseID,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PriceInfo,
		arg.StageOneOpenAt,
		arg.StageOneStartAt,
		arg.StageTwoOpenAt,
		arg.StageTwoStartAt,
		arg.Seats,
		arg.AvailableSeats,
	)
	var i LiveEvent
	err := row.Scan(
		&i.ID,
		&i.LiveHouseID,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PriceInfo,
		&i.StageOneOpenAt,
		&i.StageOneStartAt,
		&i.StageTwoOpenAt,
		&i.StageTwoStartAt,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Seats,
		&i.AvailableSeats,
	)
	return i, err
}

const getAllLiveEvents = `-- name: GetAllLiveEvents :many
SELECT
  live_houses.slug as live_house_slug,
  live_events.id,
  live_events.title, live_events.url,
  live_events.description, live_events.price_info,
  live_events.stage_one_open_at, live_events.stage_one_start_at,
  live_events.stage_two_open_at, live_events.stage_two_start_at,
  live_events.seats, live_events.available_seats,
  live_events.slug
FROM live_events
INNER JOIN live_houses ON live_events.live_house_id = live_houses.id
`

type GetAllLiveEventsRow struct {
	LiveHouseSlug   types.NullString `json:"live_house_slug"`
	ID              uuid.UUID        `json:"id"`
	Title           string           `json:"title"`
	Url             string           `json:"url"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	Seats           int32            `json:"seats"`
	AvailableSeats  int32            `json:"available_seats"`
	Slug            types.NullString `json:"slug"`
}

func (q *Queries) GetAllLiveEvents(ctx context.Context) ([]GetAllLiveEventsRow, error) {
	rows, err := q.query(ctx, q.getAllLiveEventsStmt, getAllLiveEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllLiveEventsRow
	for rows.Next() {
		var i GetAllLiveEventsRow
		if err := rows.Scan(
			&i.LiveHouseSlug,
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PriceInfo,
			&i.StageOneOpenAt,
			&i.StageOneStartAt,
			&i.StageTwoOpenAt,
			&i.StageTwoStartAt,
			&i.Seats,
			&i.AvailableSeats,
			&i.Slug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllLiveEventsByLiveHouseSlug = `-- name: GetAllLiveEventsByLiveHouseSlug :many
SELECT
  live_houses.slug as live_house_slug,
  live_events.id,
  live_events.title, live_events.url,
  live_events.description, live_events.price_info,
  live_events.stage_one_open_at, live_events.stage_one_start_at,
  live_events.stage_two_open_at, live_events.stage_two_start_at,
  live_events.seats, live_events.available_seats,
  live_events.slug
FROM live_events
INNER JOIN live_houses ON live_events.live_house_id = live_houses.id
WHERE live_houses.slug = $1
`

type GetAllLiveEventsByLiveHouseSlugRow struct {
	LiveHouseSlug   types.NullString `json:"live_house_slug"`
	ID              uuid.UUID        `json:"id"`
	Title           string           `json:"title"`
	Url             string           `json:"url"`
	Description     types.NullString `json:"description"`
	PriceInfo       types.NullString `json:"price_info"`
	StageOneOpenAt  types.NullTime   `json:"stage_one_open_at"`
	StageOneStartAt time.Time        `json:"stage_one_start_at"`
	StageTwoOpenAt  types.NullTime   `json:"stage_two_open_at"`
	StageTwoStartAt types.NullTime   `json:"stage_two_start_at"`
	Seats           int32            `json:"seats"`
	AvailableSeats  int32            `json:"available_seats"`
	Slug            types.NullString `json:"slug"`
}

func (q *Queries) GetAllLiveEventsByLiveHouseSlug(ctx context.Context, slug types.NullString) ([]GetAllLiveEventsByLiveHouseSlugRow, error) {
	rows, err := q.query(ctx, q.getAllLiveEventsByLiveHouseSlugStmt, getAllLiveEventsByLiveHouseSlug, slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllLiveEventsByLiveHouseSlugRow
	for rows.Next() {
		var i GetAllLiveEventsByLiveHouseSlugRow
		if err := rows.Scan(
			&i.LiveHouseSlug,
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PriceInfo,
			&i.StageOneOpenAt,
			&i.StageOneStartAt,
			&i.StageTwoOpenAt,
			&i.StageTwoStartAt,
			&i.Seats,
			&i.AvailableSeats,
			&i.Slug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLiveEventById = `-- name: GetLiveEventById :one
SELECT
  live_events.id, live_events.live_house_id, live_events.title, live_events.url, live_events.description, live_events.price_info, live_events.stage_one_open_at, live_events.stage_one_start_at, live_events.stage_two_open_at, live_events.stage_two_start_at, live_events.slug, live_events.created_at, live_events.updated_at, live_events.seats, live_events.available_seats,
  live_houses.name AS live_house_name,
  live_houses.slug AS live_house_slug
FROM live_events
INNER JOIN live_houses ON live_events.live_house_id = live_houses.id
WHERE live_events.id = $1
`

type GetLiveEventByIdRow struct {
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
	LiveHouseName   string           `json:"live_house_name"`
	LiveHouseSlug   types.NullString `json:"live_house_slug"`
}

func (q *Queries) GetLiveEventById(ctx context.Context, id uuid.UUID) (GetLiveEventByIdRow, error) {
	row := q.queryRow(ctx, q.getLiveEventByIdStmt, getLiveEventById, id)
	var i GetLiveEventByIdRow
	err := row.Scan(
		&i.ID,
		&i.LiveHouseID,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PriceInfo,
		&i.StageOneOpenAt,
		&i.StageOneStartAt,
		&i.StageTwoOpenAt,
		&i.StageTwoStartAt,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Seats,
		&i.AvailableSeats,
		&i.LiveHouseName,
		&i.LiveHouseSlug,
	)
	return i, err
}

const getLiveEventsByLiveHouse = `-- name: GetLiveEventsByLiveHouse :many
SELECT id, live_house_id, title, url, description, price_info, stage_one_open_at, stage_one_start_at, stage_two_open_at, stage_two_start_at, slug, created_at, updated_at, seats, available_seats FROM live_events
WHERE live_house_id = $1
LIMIT $2
OFFSET $3
`

type GetLiveEventsByLiveHouseParams struct {
	LiveHouseID uuid.UUID `json:"live_house_id"`
	Limit       int32     `json:"limit"`
	Offset      int32     `json:"offset"`
}

func (q *Queries) GetLiveEventsByLiveHouse(ctx context.Context, arg GetLiveEventsByLiveHouseParams) ([]LiveEvent, error) {
	rows, err := q.query(ctx, q.getLiveEventsByLiveHouseStmt, getLiveEventsByLiveHouse, arg.LiveHouseID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LiveEvent
	for rows.Next() {
		var i LiveEvent
		if err := rows.Scan(
			&i.ID,
			&i.LiveHouseID,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PriceInfo,
			&i.StageOneOpenAt,
			&i.StageOneStartAt,
			&i.StageTwoOpenAt,
			&i.StageTwoStartAt,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Seats,
			&i.AvailableSeats,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLiveEventAvailableSeatsById = `-- name: UpdateLiveEventAvailableSeatsById :exec
UPDATE live_events SET available_seats = $1
WHERE id = $2
`

type UpdateLiveEventAvailableSeatsByIdParams struct {
	AvailableSeats int32     `json:"available_seats"`
	ID             uuid.UUID `json:"id"`
}

func (q *Queries) UpdateLiveEventAvailableSeatsById(ctx context.Context, arg UpdateLiveEventAvailableSeatsByIdParams) error {
	_, err := q.exec(ctx, q.updateLiveEventAvailableSeatsByIdStmt, updateLiveEventAvailableSeatsById, arg.AvailableSeats, arg.ID)
	return err
}
