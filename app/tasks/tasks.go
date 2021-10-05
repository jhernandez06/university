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
	faculties := [5]string{"Ingenieria", "Medicina", "Arte", "Educacion", "Derecho"}
	courses_test := [30]string{"Electronica", "Circuitos", "Control", "Instrumentaccion", "Ortopedia", "Salud mental", "Neuro", "Farmacologia", "Arte", "Literatura", "Ceramica", "Dibujo tecnico", "Calculo 1", "Diseño experimental", "Fisica 1", "Metodos numericos", "Civil 1", "Civil 2", "Laboral 1", "Laboral 2", "Penal 1", "Penal 2"}

	for i := 0; i < 5; i++ {
		var f Fako
		fako.Fill(&f)
		dean := &models.Dean{
			FirstName:          f.FirstName, //fmt.Sprintf("Admin %v", i+1),
			LastName:           f.LastName,
			Rol:                "deans",
			IdentificationCard: fmt.Sprintf("106%v", i+1),
			CellPhoneNumber:    f.Celular}

		if err := db.Create(dean); err != nil {
			return err
		}

		faculty := models.Faculty{
			DeanID:   dean.ID,
			Number:   fmt.Sprintf("b%v", i+1),
			Location: f.Ubicacion,
			Name:     faculties[i],
		}
		if err := db.Create(&faculty); err != nil {
			return err
		}
		for i := 0; i < 4; i++ {
			var f Fako
			fako.Fill(&f)
			teacher := models.Teacher{
				FacultyID:          faculty.ID,
				FirstName:          f.FirstName,
				LastName:           f.LastName,
				IdentificationCard: fmt.Sprintf("1065%v", f.Cedula),
				JobTitle:           f.Titulo,
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
			Code:     fmt.Sprintf("00%v", i+1),
			Name:     courses_test[i],
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
	deans := models.Deans{}
	if err := q.All(&deans); err != nil {
		return err
	}
	for _, d := range deans {
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
