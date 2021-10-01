package models

type TeacherFaculty struct {
	Faculty   string `db:"faculty" rw:"r"`
	FirstName string `db:"first_name" rw:"r"`
	LastName  string `db:"last_name" rw:"r"`
	JobTitle  string `db:"job_title" rw:"r"`
}

// Facultades is not required by pop and may be deleted
type TeacherFaculties []TeacherFaculty
