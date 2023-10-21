-- name: CreateOrganization :one
INSERT INTO organizations(name)
VALUES ($1)
RETURNING *;
-- name: AddOrganizationMember :exec
INSERT INTO organization_members(user_uuid, org_uuid)
VALUES ($1, $2)
RETURNING *;
-- name: GetOrganizationByUUID :one
SELECT *
FROM organizations
WHERE uuid = $1;
-- name: GetUserOrganizations :many
SELECT o.uuid AS uuid,
    o.name AS name,
    o.created_at AS created_at
FROM users u
    JOIN organization_members om ON u.uuid = om.user_uuid
    JOIN organizations o ON om.org_uuid = o.uuid
WHERE u.uuid = $1;