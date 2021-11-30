-- name: CreateLiveHouse :one
INSERT INTO live_houses (
  name, address, slug
) VALUES (
  $1, $2, $3
) RETURNING id, name, address, slug;

-- name: GetLiveHouseById :one
SELECT * FROM live_houses
WHERE id = $1 LIMIT 1;

-- name: GetAllLiveHouses :many
SELECT * FROM live_houses
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetAllLiveHousesDefault :many
SELECT id, name, address, slug FROM live_houses
ORDER BY id;
