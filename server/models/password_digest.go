package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

//PasswordDigest for every registered user
type PasswordDigest struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	OwnerID   uuid.UUID `json:"owner_id" db:"owner_id" belongs_to:"user"`
	Digest    []byte    `json:"digest" db:"digest"`
}

// String is not required by pop and may be deleted
func (p PasswordDigest) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PasswordDigests is not required by pop and may be deleted
type PasswordDigests []PasswordDigest

// String is not required by pop and may be deleted
func (p PasswordDigests) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PasswordDigest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.BytesArePresent{Field: p.Digest, Name: "Digest"},
		&validators.UUIDIsPresent{Field: p.OwnerID, Name: "OwnerID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PasswordDigest) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PasswordDigest) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
