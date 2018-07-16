package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Horse struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Name       string    `json:"name" db:"name"`
	Age        int       `json:"age" db:"age"`
	Race       string    `json:"race" db:"race"`
	Profession string    `json:"profession" db:"profession"`
}

// String is not required by pop and may be deleted
func (h Horse) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Horses is not required by pop and may be deleted
type Horses []Horse

// String is not required by pop and may be deleted
func (h Horses) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *Horse) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: h.Name, Name: "Name"},
		&validators.IntIsPresent{Field: h.Age, Name: "Age"},
		&validators.StringIsPresent{Field: h.Race, Name: "Race"},
		&validators.StringIsPresent{Field: h.Profession, Name: "Profession"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (h *Horse) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (h *Horse) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
