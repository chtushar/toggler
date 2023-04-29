-- name: CreateUser :one
INSERT INTO users(name, email, email_verified, password, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
-- name: UpdateUser :one
-- name: DeleteUser :one