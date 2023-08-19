-- name: CreateFeatureFlag :one
INSERT INTO feature_flags(name, project_id, flag_type)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateFeatureState :one
INSERT INTO feature_states(feature_flag_id, environment_id, enabled, value)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetProjectFeatureFlags :many
SELECT ff.id,
    ff.uuid,
    ff.name,
    ff.flag_type,
    fs.enabled,
    fs.value,
    fs.updated_at
FROM feature_flags ff
    LEFT JOIN feature_states fs ON ff.id = fs.feature_flag_id
    AND fs.environment_id = $2
WHERE ff.project_id = $1;
-- name: GetFeatureFlags :many
SELECT DISTINCT ff.id AS id,
    ff.uuid AS uuid,
    ff.flag_type AS flag_type,
    ff.name AS name,
    fs.enabled AS enabled,
    fs.value AS value
FROM feature_flags ff
    JOIN project_environments pe ON ff.project_id = pe.project_id
    JOIN environments env ON pe.environment_id = env.id
    JOIN projects p ON pe.project_id = p.id
    LEFT JOIN feature_states fs ON fs.environment_id = env.id
    AND fs.feature_flag_id = ff.id
WHERE p.uuid = $1
    AND $2::text = ANY(env.api_keys);
-- name: ToggleFeatureFlag :one
UPDATE feature_states
SET enabled = NOT enabled
WHERE feature_flag_id = $1
RETURNING *;