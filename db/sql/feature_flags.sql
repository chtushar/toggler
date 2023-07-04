-- name: CreateFeatureFlag :one
INSERT INTO feature_flags(name, project_id, flag_type)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateFeatureState :one
INSERT INTO feature_states(feature_flag_id, environment_id, enabled, value)
VALUES ($1, $2, $3, $4)
RETURNING *;