-- name: CreateFeatureFlag :one
INSERT INTO feature_flags(name, project_id, flag_type)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateFeatureState :one
INSERT INTO feature_states(feature_flag_id, environment_id, enabled, value)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetProjectFeatureFlags :many
SELECT f.id AS id,
    f.uuid AS uuid,
    f.name AS name,
    f.flag_type AS flag_type,
    fs.enabled AS enabled,
    fs.value AS value,
    fs.updated_at as updated_at
FROM feature_flags f
    JOIN feature_states fs ON f.id = fs.feature_flag_id
WHERE f.project_id = $1
    AND fs.environment_id = $2;
-- name: GetFeatureFlags :many
SELECT DISTINCT ff.id AS id,
    ff.uuid AS uuid,
    ff.flag_type AS flag_type,
    ff.name AS name,
    fs.enabled AS enabled,
    fs.value AS value,
    fs.updated_at as updated_at
FROM feature_flags ff
    JOIN projects p ON ff.project_id = p.id
    JOIN environments e ON p.id = e.project_id
    JOIN feature_states fs ON ff.id = fs.feature_flag_id
    AND e.id = fs.environment_id
WHERE p.uuid = $1
    AND $2::text = ANY(env.api_keys);
-- name: ToggleFeatureFlag :one
UPDATE feature_states
SET enabled = NOT enabled
WHERE feature_flag_id = $1
RETURNING *;