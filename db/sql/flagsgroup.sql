-- name: CreateFlagsGroup :one
INSERT INTO flags_groups(name, org_id, folder_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetFolderFlagsGroup :many
SELECT *
FROM flags_groups
WHERE folder_id = $1;
-- name: GetOrgFlagsGroup :many
SELECT *
FROM flags_groups
WHERE org_id = $1;
