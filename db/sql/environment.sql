-- name: CreateEnvironment :one
INSERT INTO environments(name, project_id, api_keys)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetProjectEnvironments :many
SELECT *
FROM environments
WHERE project_id = $1;