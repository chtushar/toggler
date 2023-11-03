-- name: CreateFlagsGroupState :one
INSERT INTO flags_group_states(json, flags_group_id, environment_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetFlagsGroupState :one
SELECT *
FROM flags_group_states
WHERE flags_group_id = $1 AND environment_id = $2;
-- name: GetFlagsGroupStateByUUID :one
SELECT *
FROM flags_group_states
WHERE uuid = $1;
-- name: UpdateFlagGroupsStateJSON :one
UPDATE flags_group_states
SET json = $1
WHERE flags_group_id = $2 AND environment_id = $3
RETURNING *;
