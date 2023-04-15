package models

import "time"

type Project struct {
    ID        int64      `json:"id"`
    Name      string     `json:"name"`
    OwnerID   int64      `json:"owner_id"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

type User struct {
    ID            int64      `json:"id"`
    Name          string     `json:"name"`
    Password      string     `json:"-"`
    Email         string     `json:"email"`
    EmailVerified bool       `json:"email_verified"`
    CreatedAt     *time.Time `json:"created_at"`
    UpdatedAt     *time.Time `json:"updated_at"`
}

type FeatureFlag struct {
    ID      int64  `json:"id"`
    UUID    string `json:"uuid"`
    ProjectID  int64  `json:"project_id"`
    Type    int    `json:"type"`
}

type ProjectMember struct {
    UserID    int64      `json:"user_id"`
    ProjectID    int64      `json:"project_id"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

type FFResolutionBoolean struct {
    ID      int64  `json:"id"`
    Key     string `json:"key"`
    FFID    int64  `json:"ff_id"`
    Enabled bool   `json:"enabled"`
}

type FeatureFlagType struct {
    ID   int64  `json:"id"`
    Type string `json:"type"`
}
