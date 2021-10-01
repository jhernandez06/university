package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Decano is used by pop to map your decanoes database table to your go code.
type Dean struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	FirstName          string    `json:"first_name" db:"first_name"`
	LastName           string    `json:"last_name" db:"last_name"`
	IdentificationCard string    `json:"identification_card" db:"identification_card"`
	Rol                string    `json:"rol" db:"rol"`
	CellPhoneNumber    string    `json:"cell_phone_number" db:"cell_phone_number"`
	Faculty            Faculty   `has_one:"faculty"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (d Dean) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Decanoes is not required by pop and may be deleted
type Deans []Dean

// String is not required by pop and may be deleted
func (d Deans) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *Dean) Validate(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: d.FirstName, Name: "First Name"},
		&validators.StringIsPresent{Field: d.LastName, Name: "Last Name"},
		&validators.StringIsPresent{Field: d.IdentificationCard, Name: "Identification card"},
		&validators.StringIsPresent{Field: d.Rol, Name: "Rol"},
		&validators.StringIsPresent{Field: d.CellPhoneNumber, Name: "Cell phone number"},
	)
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *Dean) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *Dean) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
