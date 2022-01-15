-- name: CreateUserPoints :one
INSERT INTO user_points (
  user_id, points, description, order_type, order_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserPoints :one
SELECT COALESCE(SUM(points), 0) FROM user_points WHERE user_id = $1;

-- name: DeleteUserPointsByOrderID :exec
DELETE FROM user_points WHERE order_id = $1;
