-- name: CreateActiveUser :one
INSERT INTO users (name, email, password, email_verified, active)
VALUES ($1, $2, $3, TRUE, TRUE)
RETURNING *;