// Code generated by sqlc. DO NOT EDIT.
// source: live_house.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createLiveHouse = `-- name: CreateLiveHouse :one
INSERT INTO live_houses (
  name, address, slug
) VALUES (
  $1, $2, $3
) RETURNING id, name, address, slug, created_at, updated_at
`

type CreateLiveHouseParams struct {
	Name    string         `json:"name"`
	Address sql.NullString `json:"address"`
	Slug    sql.NullString `json:"slug"`
}

func (q *Queries) CreateLiveHouse(ctx context.Context, arg CreateLiveHouseParams) (LiveHouse, error) {
	row := q.queryRow(ctx, q.createLiveHouseStmt, createLiveHouse, arg.Name, arg.Address, arg.Slug)
	var i LiveHouse
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllLiveHouses = `-- name: GetAllLiveHouses :many
SELECT id, name, address, slug, created_at, updated_at FROM live_houses
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetAllLiveHousesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllLiveHouses(ctx context.Context, arg GetAllLiveHousesParams) ([]LiveHouse, error) {
	rows, err := q.query(ctx, q.getAllLiveHousesStmt, getAllLiveHouses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LiveHouse
	for rows.Next() {
		var i LiveHouse
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getLiveHouseById = `-- name: GetLiveHouseById :one
SELECT id, name, address, slug, created_at, updated_at FROM live_houses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLiveHouseById(ctx context.Context, id uuid.UUID) (LiveHouse, error) {
	row := q.queryRow(ctx, q.getLiveHouseByIdStmt, getLiveHouseById, id)
	var i LiveHouse
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
