-- name: CreateProject :one
INSERT INTO projects(name, owner_id, org_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: AddProjectMember :exec
INSERT INTO project_members(user_id, project_id)
VALUES ($1, $2);
-- name: GetProject :one
SELECT *
FROM projects
WHERE id = $1;
-- name: GetUserProjects :many
SELECT p.*
FROM projects p
    INNER JOIN project_members pm ON pm.project_id = p.id
WHERE pm.user_id = $1;