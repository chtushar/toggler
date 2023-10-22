-- name: CreateEnvironment :exec
INSERT INTO environments(name, color, org_id)
VALUES ($1, $2, $3);
-- name: GetOrganizationEnvironments :many
SELECT *
FROM environments
WHERE org_id = $1;
-- name: UpdateEnvironmentName :exec
UPDATE environments
SET name = $2
WHERE uuid = $1;
-- name: UpdateEnvironmentColor :exec
UPDATE environments
SET color = $2
WHERE uuid = $1;