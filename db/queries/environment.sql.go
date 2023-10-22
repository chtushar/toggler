// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: environment.sql

package queries

import (
	"context"

	"github.com/google/uuid"
)

const createEnvironment = `-- name: CreateEnvironment :one
INSERT INTO environments(name, color, org_uuid)
VALUES ($1, $2, $3)
RETURNING uuid, id, name, color, org_uuid, created_at
`

type CreateEnvironmentParams struct {
	Name    string     `json:"name"`
	Color   *string    `json:"color"`
	OrgUuid *uuid.UUID `json:"org_uuid"`
}

func (q *Queries) CreateEnvironment(ctx context.Context, arg CreateEnvironmentParams) (Environment, error) {
	row := q.db.QueryRow(ctx, createEnvironment, arg.Name, arg.Color, arg.OrgUuid)
	var i Environment
	err := row.Scan(
		&i.Uuid,
		&i.ID,
		&i.Name,
		&i.Color,
		&i.OrgUuid,
		&i.CreatedAt,
	)
	return i, err
}