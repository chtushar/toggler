-- name: CreateEnvironment :one
INSERT INTO environments(name, color, id)
VALUES ($1, $2, $3)
RETURNING *;