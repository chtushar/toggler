-- name: CreateUser :one
INSERT INTO users(name, email, email_verified, password)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: CountUsers :one
SELECT COUNT(*)
FROM users;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetAllUsers :many
SELECT *
FROM users;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
-- name: UpdateUser :one
UPDATE users
SET name = $1,
    email = $2,
    email_verified = $3
WHERE id = $4
RETURNING *;
-- name: UpdateUserPassword :one
UPDATE users
SET password = $1
WHERE id = $2
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;