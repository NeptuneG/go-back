-- name: CreateUser :one
INSERT INTO users (
  email, encrypted_password
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;
