package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Course is used by pop to map your courses database table to your go code.
type Course struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Codigo    string    `json:"codigo" db:"codigo"`
	Nombre    string    `json:"nombre" db:"nombre"`
	Creditos  int       `json:"creditos" db:"creditos"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Course) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Courses is not required by pop and may be deleted
type Courses []Course

// String is not required by pop and may be deleted
func (c Courses) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Course) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Codigo, Name: "Codigo"},
		&validators.StringIsPresent{Field: c.Nombre, Name: "Nombre"},
		&validators.IntIsPresent{Field: c.Creditos, Name: "Creditos"},
	)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Course) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Course) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
