-- name: CreateUserPoints :one
INSERT INTO user_points (
  user_id, points, description, order_type, order_id, tx_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserPoints :one
SELECT COALESCE(SUM(points), 0) FROM user_points WHERE user_id = $1;

-- name: DeleteUserPointsByTxID :exec
DELETE FROM user_points WHERE tx_id = $1;
