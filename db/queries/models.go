// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package queries

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type FeatureFlagType string

const (
	FeatureFlagTypeBoolean FeatureFlagType = "boolean"
)

func (e *FeatureFlagType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FeatureFlagType(s)
	case string:
		*e = FeatureFlagType(s)
	default:
		return fmt.Errorf("unsupported scan type for FeatureFlagType: %T", src)
	}
	return nil
}

type NullFeatureFlagType struct {
	FeatureFlagType FeatureFlagType
	Valid           bool // Valid is true if FeatureFlagType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFeatureFlagType) Scan(value interface{}) error {
	if value == nil {
		ns.FeatureFlagType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FeatureFlagType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFeatureFlagType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FeatureFlagType), nil
}

type UserRole string

const (
	UserRoleMember UserRole = "member"
	UserRoleAdmin  UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Environment struct {
	ID        int32         `json:"id"`
	Name      string        `json:"name"`
	ApiKeys   []string      `json:"api_keys"`
	Uuid      uuid.NullUUID `json:"uuid"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type FeatureFlag struct {
	ID        int32           `json:"id"`
	ProjectID int64           `json:"project_id"`
	Uuid      uuid.NullUUID   `json:"uuid"`
	FlagType  FeatureFlagType `json:"flag_type"`
	Name      string          `json:"name"`
}

type FeatureState struct {
	ID            int32         `json:"id"`
	Uuid          uuid.NullUUID `json:"uuid"`
	EnvironmentID int64         `json:"environment_id"`
	FeatureFlagID int64         `json:"feature_flag_id"`
	Enabled       bool          `json:"enabled"`
	Value         pgtype.JSONB  `json:"value"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type Organization struct {
	ID        int32         `json:"id"`
	Uuid      uuid.NullUUID `json:"uuid"`
	Name      string        `json:"name"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type OrganizationMember struct {
	UserID    int64     `json:"user_id"`
	OrgID     int64     `json:"org_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrganizationOnboarding struct {
	OrgID         int64        `json:"org_id"`
	CreateProject sql.NullBool `json:"create_project"`
}

type Project struct {
	ID        int32         `json:"id"`
	Name      string        `json:"name"`
	Uuid      uuid.NullUUID `json:"uuid"`
	OrgID     int64         `json:"org_id"`
	OwnerID   int64         `json:"owner_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type ProjectEnvironment struct {
	ProjectID     int64     `json:"project_id"`
	EnvironmentID int64     `json:"environment_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProjectMember struct {
	UserID    int64     `json:"user_id"`
	ProjectID int64     `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID            int32          `json:"id"`
	Name          string         `json:"name"`
	Uuid          uuid.NullUUID  `json:"uuid"`
	Password      string         `json:"password"`
	Email         sql.NullString `json:"email"`
	EmailVerified bool           `json:"email_verified"`
	Role          UserRole       `json:"role"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
