package models

import "time"

type Project struct {
	ID        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	OwnerID   int64      `json:"owner_id" db:"owner_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	ID            int64      `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Password      string     `json:"-" db:"password"`
	Email         string     `json:"email" db:"email"`
	EmailVerified bool       `json:"email_verified" db:"email_verified"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	Role          string     `json:"role" db:"role"`
}

type FeatureFlag struct {
	ID        int64  `json:"id" db:"id"`
	UUID      string `json:"uuid" db:"uuid"`
	ProjectID int64  `json:"project_id" db:"team_id"`
	Type      int    `json:"type" db:"type"`
	Name      string `json:"name" db:"name"`
}

type ProjectMember struct {
	UserID    int64      `json:"user_id" db:"user_id"`
	ProjectID int64      `json:"project_id" db:"project_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
