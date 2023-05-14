-- name: CreateProject :one
INSERT INTO projects(name, owner_id)
VALUES ($1, $2)
RETURNING *;

-- name: AddProjectMember :exec
INSERT INTO project_members(user_id, project_id)
VALUES ($1, $2);

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;
