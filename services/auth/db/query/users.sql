-- name: CreateUser :one
INSERT INTO users (
  email, encrypted_password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE users.id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE users.email = $1;

