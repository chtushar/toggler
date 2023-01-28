package models

import "time"

type Base struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type User struct {
	Base

	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `db:"password"`
}

type TeamMember struct {
	Base

	TeamID int `json:"team_id" db:"team_id"`
	UserID int `json:"user_id" db:"user_id"`
}

type Team struct {
	Base

	Name  string `json:"name" db:"name"`
	Owner int    `json:"owner" db:"owner"`
}

type FeatureFlagTypes struct {
	ID   int    `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
}

type FeatureFlag struct {
	Base

	Name   string `json:"name" db:"name"`
	Type   int    `json:"type" db:"type"`
	TeamID int    `json:"team_id" db:"team_id"`
}
