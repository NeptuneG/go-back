-- name: CreateUser :one
INSERT INTO users (
  email, encrypted_password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserByID :one
SELECT
  users.id,
  users.email,
  SUM(user_points.points) as points
FROM users
INNER JOIN user_points ON users.id = user_points.user_id
WHERE users.id = $1
GROUP BY users.id;

-- name: GetUserByEmail :one
SELECT
  users.id,
  users.email,
  SUM(user_points.points) as points
FROM users
INNER JOIN user_points ON users.id = user_points.user_id
WHERE users.email = $1
GROUP BY users.id;

-- name: IsUserExist :one
SELECT EXISTS(SELECT users.id from users where users.id = $1);
