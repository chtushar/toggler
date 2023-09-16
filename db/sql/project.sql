-- name: CreateProject :one
INSERT INTO projects(name, owner_id, org_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetProject :one
SELECT *
FROM projects
WHERE id = $1;
-- name: GetUserOrgProjects :many
SELECT id,
    name,
    uuid,
    org_id
FROM projects p
WHERE p.org_id = $2
    AND EXISTS (
        SELECT *
        FROM organization_members
        WHERE user_id = $1
            AND org_id = $2
    );