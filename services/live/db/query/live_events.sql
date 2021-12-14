-- name: CreateLiveEvent :one
INSERT INTO live_events (
  live_house_id, title, url,
  description, price_info,
  stage_one_open_at, stage_one_start_at,
  stage_two_open_at, stage_two_start_at,
  available_seats
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetLiveEventById :one
SELECT * FROM live_events
WHERE id = $1 LIMIT 1;

-- name: GetLiveEventsByLiveHouse :many
SELECT * FROM live_events
WHERE live_house_id = $1
LIMIT $2
OFFSET $3;

-- name: GetAllLiveEvents :many
SELECT
  live_houses.slug as live_house_slug,
  live_events.id,
  live_events.title, live_events.url,
  live_events.description, live_events.price_info,
  live_events.stage_one_open_at, live_events.stage_one_start_at,
  live_events.stage_two_open_at, live_events.stage_two_start_at,
  live_events.available_seats,
  live_events.slug
FROM live_events
INNER JOIN live_houses ON live_events.live_house_id = live_houses.id;

-- name: GetAllLiveEventsByLiveHouseSlug :many
SELECT
  live_houses.slug as live_house_slug,
  live_events.id,
  live_events.title, live_events.url,
  live_events.description, live_events.price_info,
  live_events.stage_one_open_at, live_events.stage_one_start_at,
  live_events.stage_two_open_at, live_events.stage_two_start_at,
  live_events.available_seats,
  live_events.slug
FROM live_events
INNER JOIN live_houses ON live_events.live_house_id = live_houses.id
WHERE live_houses.slug = $1;

-- name: UpdateLiveEventAvailableSeatsById :exec
UPDATE live_events SET available_seats = $1
WHERE id = $2;
