-- name: CreateEnvironment :one
INSERT INTO environments(name)
VALUES ($1)
RETURNING *;
-- name: AddProjectEnvironment :exec
INSERT INTO project_environments(project_id, environment_id)
VALUES ($1, $2);
-- name: CreateProdAndDevEnvironments :many
INSERT INTO environments(name, api_keys)
VALUES ('production', $1),
    ('development', $2)
RETURNING *;
-- name: AddProdAndDevProjectEnvironments :exec
INSERT INTO project_environments(project_id, environment_id)
VALUES ($1, $2),
    ($1, $3);
-- name: GetProjectEnvironments :many
SELECT e.*
FROM environments e
    INNER JOIN project_environments pe ON pe.environment_id = e.id
WHERE pe.project_id = $1;