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
type Facultad struct {
	ID           uuid.UUID `json:"id" db:"id"`
	DecanoID     uuid.UUID `json:"decano_id" db:"decano_id"`
	Nombre       string    `json:"nombre" db:"nombre"`
	Numero       string    `json:"numero" db:"numero"`
	Ubicacion    string    `json:"ubicacion" db:"ubicacion"`
	Decano       *Decano   `json:"decano,omitempty" belongs_to:"decano"`
	DecanoNombre string    `json:"-" db:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (f Facultad) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Facultades is not required by pop and may be deleted
type Facultades []Facultad

// String is not required by pop and may be deleted
func (f Facultades) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Facultad) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: f.DecanoID, Name: "DecanoID"},
		&validators.StringIsPresent{Field: f.Nombre, Name: "Nombre"},
		&validators.StringIsPresent{Field: f.Numero, Name: "Numero"},
		&validators.StringIsPresent{Field: f.Ubicacion, Name: "Ubicacion"},
	)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Facultad) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Facultad) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
