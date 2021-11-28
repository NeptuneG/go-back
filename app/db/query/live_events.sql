-- name: CreateLiveEvent :one
INSERT INTO live_events (
  live_house_id, title, url,
  description, price_info,
  stage_one_open_at, stage_one_start_at,
  stage_two_open_at, stage_two_start_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetLiveEventById :one
SELECT * FROM live_events
WHERE id = $1 LIMIT 1;

-- name: GetLiveEventsByLiveHouse :many
SELECT * FROM live_events
WHERE live_house_id = $1
LIMIT $2
OFFSET $3;
