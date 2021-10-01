package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Facultade is used by pop to map your facultades database table to your go code.
type Faculty struct {
	ID        uuid.UUID `json:"id" db:"id"`
	DeanID    uuid.UUID `json:"dean_id" db:"dean_id"`
	Name      string    `json:"name" db:"name"`
	Number    string    `json:"number" db:"number"`
	Location  string    `json:"location" db:"location"`
	Dean      *Dean     `json:"dean,omitempty" belongs_to:"dean"`
	DeanName  string    `json:"-" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (f Faculty) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Facultades is not required by pop and may be deleted
type Faculties []Faculty

// String is not required by pop and may be deleted
func (f Faculties) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Faculty) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: f.DeanID, Name: "DeanID"},
		&validators.StringIsPresent{Field: f.Name, Name: "Name"},
		&validators.StringIsPresent{Field: f.Number, Name: "Number"},
		&validators.StringIsPresent{Field: f.Location, Name: "Location"},
	)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Faculty) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Faculty) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
