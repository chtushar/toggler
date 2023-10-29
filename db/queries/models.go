// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package queries

import (
	"time"

	"github.com/jackc/pgtype"
)

type ApiKey struct {
	Uuid           string     `json:"uuid"`
	ID             *int32     `json:"-"`
	Name           string     `json:"name"`
	ApiKey         string     `json:"api_key"`
	AllowedDomains []string   `json:"allowed_domains"`
	OrgID          *int32     `json:"-"`
	UserID         *int32     `json:"-"`
	CreatedAt      *time.Time `json:"created_at"`
}

type Environment struct {
	Uuid      string     `json:"uuid"`
	ID        *int32     `json:"-"`
	Name      string     `json:"name"`
	Color     *string    `json:"color"`
	OrgID     *int32     `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
}

type FlagsGroup struct {
	Uuid           string     `json:"uuid"`
	ID             *int32     `json:"-"`
	Name           string     `json:"name"`
	OrgID          *int32     `json:"-"`
	FolderID       *int32     `json:"-"`
	CurrentVersion *int32     `json:"current_version"`
	CreatedAt      *time.Time `json:"created_at"`
}

type FlagsGroupState struct {
	Uuid          string       `json:"uuid"`
	ID            *int32       `json:"-"`
	FlagsGroupID  *int32       `json:"-"`
	Version       *int32       `json:"version"`
	Json          pgtype.JSONB `json:"json"`
	EnvironmentID *int32       `json:"-"`
	CreatedAt     *time.Time   `json:"created_at"`
}

type Folder struct {
	Uuid      string     `json:"uuid"`
	ID        *int32     `json:"-"`
	Name      string     `json:"name"`
	OrgID     *int32     `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
}

type Organization struct {
	Uuid      string     `json:"uuid"`
	ID        *int32     `json:"-"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type OrganizationMember struct {
	UserID *int32 `json:"-"`
	OrgID  *int32 `json:"-"`
}

type User struct {
	Uuid          string     `json:"uuid"`
	ID            *int32     `json:"-"`
	Name          string     `json:"name"`
	Email         string     `json:"email"`
	Password      string     `json:"-"`
	EmailVerified *bool      `json:"email_verified"`
	Active        *bool      `json:"active"`
	CreatedAt     *time.Time `json:"created_at"`
}
