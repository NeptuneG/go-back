-- name: CreateLiveEventOrder :one
INSERT INTO live_event_orders (
  user_id, live_event_id, price, user_points
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateLiveEventOrderState :exec
UPDATE live_event_orders SET state = $1
WHERE id = $2;
