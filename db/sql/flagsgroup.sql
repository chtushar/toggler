-- name: CreateFlagsGroup :one
INSERT INTO flags_groups(name, org_id, folder_id, current_version)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateFlagsGroupCurrentVersion :exec
UPDATE flags_groups
SET current_version = $1
WHERE uuid = $2;
-- name: GetFolderFlagsGroup :many
SELECT *
FROM flags_groups
WHERE folder_id = $1;
-- name: GetOrgFlagsGroup :many
SELECT *
FROM flags_groups
WHERE org_id = $1;