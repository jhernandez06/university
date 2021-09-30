package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Teacher is used by pop to map your teachers database table to your go code.
type Teacher struct {
	ID         uuid.UUID `json:"id" db:"id"`
	FacultadID uuid.UUID `json:"facultad_id" db:"facultad_id"`
	Nombre     string    `json:"nombre" db:"nombre"`
	Apellido   string    `json:"apellido" db:"apellido"`
	Cedula     string    `json:"cedula" db:"cedula"`
	Titulo     string    `json:"titulo" db:"titulo"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Teacher) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Teachers is not required by pop and may be deleted
type Teachers []Teacher

// String is not required by pop and may be deleted
func (t Teachers) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Teacher) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Nombre, Name: "Nombre"},
		&validators.StringIsPresent{Field: t.Apellido, Name: "Apellido"},
		&validators.StringIsPresent{Field: t.Cedula, Name: "Cedula"},
		&validators.StringIsPresent{Field: t.Titulo, Name: "Titulo"},
	)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Teacher) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Teacher) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
