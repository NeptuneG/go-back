-- name: CreateUserOrder :one
INSERT INTO user_orders (
  user_id, live_event_id
) VALUES (
  $1, $2
) RETURNING *;
