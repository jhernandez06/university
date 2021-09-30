package models

type ViewTeacherCourse struct {
	Course   string `db:"course" rw:"r"`
	Facultad string `db:"facultad" rw:"r"`
	Nombre   string `db:"nombre" rw:"r"`
	Apellido string `db:"apellido" rw:"r"`
	Titulo   string `db:"titulo" rw:"r"`
}

// Facultades is not required by pop and may be deleted
type ViewTeacherCourses []ViewTeacherCourse
