-- name: CreateFlagsGroupState :one
INSERT INTO flags_group_states(flags_group_id)
VALUES ($1)
RETURNING *;
-- name: GetFlagsGroupState :one
SELECT *
FROM flags_group_states
WHERE flags_group_id = $1;
-- name: GetFlagsGroupStateByUUID :one
SELECT *
FROM flags_group_states
WHERE uuid = $1;
-- name: UpdateFlagGroupsStateJS :one
UPDATE flags_group_states
SET js = $1
WHERE flags_group_id = $2
RETURNING *;