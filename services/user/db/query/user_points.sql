-- name: CreateUserPoints :one
INSERT INTO user_points (
  user_id, points, description, order_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserPoints :one
SELECT SUM(points) FROM user_points WHERE user_id = $1;

-- name: DeleteUserPointsByOrderID :exec
DELETE FROM user_points WHERE order_id = $1;
