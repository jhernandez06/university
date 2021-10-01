package models

type ViewTeacherCourse struct {
	Course    string `db:"course" rw:"r"`
	Faculty   string `db:"faculty" rw:"r"`
	FirstName string `db:"first_name" rw:"r"`
	LastName  string `db:"last_name" rw:"r"`
	JobTitle  string `db:"job_title" rw:"r"`
}

type ViewTeacherCourses []ViewTeacherCourse
