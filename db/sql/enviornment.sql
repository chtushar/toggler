-- name: CreateEnvironment :one
INSERT INTO environments(name)
VALUES ($1)
RETURNING *;

-- name: AddProjectEnvironment :exec
INSERT INTO project_enviornments(project_id, environment_id)
VALUES ($1, $2);

-- name: CreateProdAndDevEnvironments :many
INSERT INTO environments(name)
VALUES ('production'), ('development')
RETURNING *;

-- name: AddProdAndDevProjectEnviornments :exec
INSERT INTO project_enviornments(project_id, environment_id)
VALUES ($1, $2), ($1, $3);
