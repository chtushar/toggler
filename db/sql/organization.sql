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