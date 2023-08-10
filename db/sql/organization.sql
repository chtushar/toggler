-- name: CreateOrganization :one
INSERT INTO organizations(name)
VALUES ($1)
RETURNING *;
-- name: AddOrganizationMember :exec
INSERT INTO organization_members(user_id, org_id)
VALUES ($1, $2);
-- name: GetOrganization :one
SELECT *
FROM organizations
WHERE id = $1;
-- name: GetUserOrganizations :many
SELECT o.*
FROM organizations o
    INNER JOIN organization_members om ON om.org_id = o.id
WHERE om.user_id = $1;
-- name: DoesUserBelongToOrg :one
SELECT EXISTS (
        SELECT 1
        FROM organization_members
        WHERE user_id = $1
            AND org_id = $2
    ) AS user_belongs_to_organization;
-- name: UpdateOrganization :one
UPDATE organizations
set name = $2
WHERE id = $1
RETURNING *;