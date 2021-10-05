package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Course is used by pop to map your courses database table to your go code.
type Course struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Creditos  int       `json:"creditos" db:"creditos"`
	KeyWord   string    `json:"-" db:"-"`
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
		&validators.StringIsPresent{Field: c.Code, Name: "Code"},
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&IntIsPresent{Name: "Creditos", Field: c.Creditos},
	)
}

type IntIsPresent struct {
	Name    string
	Field   int
	Message string
}

// IsValid adds an error if the field equals 0.
func (v *IntIsPresent) IsValid(errors *validate.Errors) {
	if v.Field > 0 {
		return
	}

	errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s The credits must be greater than 0", v.Name))
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
