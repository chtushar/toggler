-- name: CreateAPIKey :one
INSERT INTO api_keys(name, api_key, allowed_domains, org_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetOrganizationAPIKeys :many
SELECT *
FROM api_keys
WHERE org_id = $1;
-- name: DeleteAPIKey :exec
DELETE FROM api_keys
WHERE uuid = $1;