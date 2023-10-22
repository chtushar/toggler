-- name: CreateEnvironment :exec
INSERT INTO environments(name, color, org_id)
VALUES ($1, $2, $3);
-- name: GetOrganizationEnvironments :many
SELECT *
FROM environments
WHERE org_id = $1;