-- name: CreateFlagsGroupState :one
INSERT INTO flags_group_states(json, flags_group_id, environment_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetFlagsGroupStateByID :one
SELECT *
FROM flags_group_states
WHERE id = $1;
-- name: GetFlagsGroupStateByUUID :one
SELECT *
FROM flags_group_states
WHERE uuid = $1;
