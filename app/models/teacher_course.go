package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
)
// TeacherCourse is used by pop to map your teacher_courses database table to your go code.
type TeacherCourse struct {
    ID uuid.UUID `json:"id" db:"id"`
    CourseID uuid.UUID `json:"course_id" db:"course_id"`
    TeacherID uuid.UUID `json:"teacher_id" db:"teacher_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t TeacherCourse) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// TeacherCourses is not required by pop and may be deleted
type TeacherCourses []TeacherCourse

// String is not required by pop and may be deleted
func (t TeacherCourses) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *TeacherCourse) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *TeacherCourse) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *TeacherCourse) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
