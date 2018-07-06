package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type UserPermission struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id"`
	PermissionID uuid.UUID  `json:"permission_id" db:"permission_id"`
	User         User       `db:"-" belongs_to:"users"`
	Perm         Permission `db:"-" belongs_to:"permissions"`
}

// String is not required by pop and may be deleted
func (u UserPermission) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserPermissions is not required by pop and may be deleted
type UserPermissions []UserPermission

// String is not required by pop and may be deleted
func (u UserPermissions) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserPermission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserPermission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserPermission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
