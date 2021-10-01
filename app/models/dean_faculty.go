package models

type DeanFaculty struct {
	FirstName   string `db:"first_name" rw:"r"`
	LastName    string `db:"last_name" rw:"r"`
	FacultyName string `db:"faculty_name" rw:"r"`
	Location    string `db:"location" rw:"r"`
	Number      string `db:"number" rw:"r"`
}

// Facultades is not required by pop and may be deleted
type DeanFaculties []DeanFaculty
