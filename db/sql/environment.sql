-- name: CreateEnvironment :exec
INSERT INTO environments(name, color, org_id)
VALUES ($1, $2, $3);