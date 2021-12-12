-- name: CreateUserPoints :one
INSERT INTO user_points (
  user_id, points, description
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUserPoints :one
SELECT SUM(points) FROM user_points WHERE user_id = $1;
