// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: organization.sql

package queries

import (
	"context"
)

const addOrganizationMember = `-- name: AddOrganizationMember :exec
INSERT INTO organization_members(user_id, org_id)
VALUES ($1, $2)
`

type AddOrganizationMemberParams struct {
	UserID int64 `json:"user_id"`
	OrgID  int64 `json:"org_id"`
}

func (q *Queries) AddOrganizationMember(ctx context.Context, arg AddOrganizationMemberParams) error {
	_, err := q.db.Exec(ctx, addOrganizationMember, arg.UserID, arg.OrgID)
	return err
}

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations(name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateOrganization(ctx context.Context, name string) (Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, name)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrganization = `-- name: GetOrganization :one
SELECT id, name, created_at, updated_at
FROM organizations
WHERE id = $1
`

func (q *Queries) GetOrganization(ctx context.Context, id int32) (Organization, error) {
	row := q.db.QueryRow(ctx, getOrganization, id)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserOrganizations = `-- name: GetUserOrganizations :many
SELECT o.id, o.name, o.created_at, o.updated_at
FROM organizations o
    INNER JOIN organization_members om ON om.org_id = o.id
WHERE om.user_id = $1
`

func (q *Queries) GetUserOrganizations(ctx context.Context, userID int64) ([]Organization, error) {
	rows, err := q.db.Query(ctx, getUserOrganizations, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
