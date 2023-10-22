-- name: CreateEnvironment :one
INSERT INTO environments(name, color, org_uuid)
VALUES ($1, $2, $3)
RETURNING *;