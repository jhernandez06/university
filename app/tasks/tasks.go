package tasks

import (
	"fmt"
	"university/app"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/grift/grift"
	"github.com/wawandco/fako"
)

type Fako struct {
	FirstName string `fako:"first_name"`
	LastName  string `fako:"last_name"`
	Email     string `fako:"email_address"`
	Titulo    string `fako:"job_title"`
	Celular   string `fako:"phone"`
	Cedula    string `fako:"digits"`
	Ubicacion string `fako:"city"`
	Codigo    string `fako:"characters"`
}

// Init the tasks with some common tasks that come from
// grift
func init() {
	buffalo.Grifts(app.New())
}

// task for create 5 teacher, 5 faculties,
var _ = grift.Add("create", func(c *grift.Context) error {
	db := models.DB()
	facultades := [5]string{"Ingenieria", "Medicina", "Arte", "Educacion", "Derecho"}
	asignaturas := [30]string{"Electronica", "Circuitos", "Control", "Instrumentaccion", "Ortopedia", "Salud mental", "Neuro", "Farmacologia", "Arte", "Literatura", "Ceramica", "Dibujo tecnico", "Calculo 1", "Diseño experimental", "Fisica 1", "Metodos numericos", "Civil 1", "Civil 2", "Laboral 1", "Laboral 2", "Penal 1", "Penal 2"}

	for i := 0; i < 5; i++ {
		var f Fako
		fako.Fill(&f)
		decano := &models.Decano{
			Nombre:   f.FirstName, //fmt.Sprintf("Admin %v", i+1),
			Apellido: f.LastName,
			Rol:      "decano",
			Cedula:   fmt.Sprintf("106%v", i+1),
			Celular:  f.Celular}

		if err := db.Create(decano); err != nil {
			return err
		}

		facultad := models.Facultad{
			DecanoID:  decano.ID,
			Numero:    fmt.Sprintf("b%v", i+1),
			Ubicacion: f.Ubicacion,
			Nombre:    facultades[i],
		}
		if err := db.Create(&facultad); err != nil {
			return err
		}
		for i := 0; i < 4; i++ {
			var f Fako
			fako.Fill(&f)
			teacher := models.Teacher{
				FacultadID: facultad.ID,
				Nombre:     f.FirstName,
				Apellido:   f.LastName,
				Cedula:     fmt.Sprintf("1065%v", f.Cedula),
				Titulo:     f.Titulo,
			}
			if err := db.Create(&teacher); err != nil {
				return err
			}
		}
	}
	for i := 0; i < 20; i++ {
		var f Fako
		fako.Fill(&f)
		course := models.Course{
			Codigo:   fmt.Sprintf("AB%v", i+1),
			Nombre:   asignaturas[i],
			Creditos: 4,
		}
		if err := db.Create(&course); err != nil {
			return err
		}
	}
	teachers := models.Teachers{}
	courses := models.Courses{}
	if err := db.All(&courses); err != nil {
		return err
	}
	if err := db.All(&teachers); err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		teacher_course := models.TeacherCourse{
			CourseID:  courses[i].ID,
			TeacherID: teachers[i].ID,
		}
		if err := db.Create(&teacher_course); err != nil {
			return err
		}
	}

	return nil
})

var _ = grift.Add("delete", func(c *grift.Context) error {
	db := models.DB()
	q := db.Q()
	decanos := models.Decanoes{}
	if err := q.All(&decanos); err != nil {
		return err
	}
	for _, d := range decanos {
		if err := db.Destroy(&d); err != nil {
			return err
		}
	}

	teachers := models.Teachers{}
	if err := q.All(&teachers); err != nil {
		return err
	}
	for _, d := range teachers {
		if err := db.Destroy(&d); err != nil {
			return err
		}
	}

	courses := models.Courses{}
	if err := q.All(&courses); err != nil {
		return err
	}
	for _, d := range courses {
		if err := db.Destroy(&d); err != nil {
			return err
		}
	}

	return nil
})

// var _ = grift.Add("facultades", func(c *grift.Context) error {
// 	db := models.DB()

// 	return nil
// })
// var _ = grift.Add("deleteFacultades", func(c *grift.Context) error {
// 	db := models.DB()

// 	return nil
// })
