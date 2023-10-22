-- name: CreateFolder :one
INSERT INTO folders(name, org_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetOrgFolders :many
SELECT *
FROM folders
WHERE org_id = $1;
-- name: GetFolderByUUID :one
SELECT *
FROM folders
WHERE org_id = $1;
-- name: UpdateFolderName :exec
UPDATE folders
SET name = $1
WHERE uuid = $2;