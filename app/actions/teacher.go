package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func NewTeacher(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculties := models.Faculties{}
	FacultiesList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("name").All(&faculties); err != nil {
		return err
	}
	for _, f := range faculties {
		oneFaculty := map[string]interface{}{
			f.Name: uuid.FromStringOrNil(f.ID.String()),
		}
		FacultiesList = append(FacultiesList, oneFaculty)
	}
	c.Set("FacultiesList", FacultiesList)
	c.Set("teacher", models.Teacher{})
	return c.Render(http.StatusOK, r.HTML("teacher/new.plush.html"))
}

func CreateTeacher(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculties := models.Faculties{}
	teacher := models.Teacher{}
	FacultiesList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("name").All(&faculties); err != nil {
		return err
	}
	for _, f := range faculties {
		oneFaculty := map[string]interface{}{
			f.Name: uuid.FromStringOrNil(f.ID.String()),
		}
		FacultiesList = append(FacultiesList, oneFaculty)
	}
	if err := c.Bind(&teacher); err != nil {
		return err
	}

	verrs := teacher.Validate(tx)
	if verrs.HasAny() {
		c.Set("teacher", teacher)
		c.Set("FacultiesList", FacultiesList)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("teacher/new.plush.html"))
	}

	if err := tx.Create(&teacher); err != nil {
		return err
	}
	c.Flash().Add("success", "teacher create succesfully")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListTeacher(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	teachers := models.Teachers{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("first_name").All(&teachers); err != nil {
		return err
	}

	c.Set("teachers", teachers)
	c.Set("paginatorT", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("teacher/list.plush.html"))
}

func DeleteTeacher(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	teacher := models.Teacher{}
	teacherID := c.Param("teacher_id")
	if err := tx.Find(&teacher, teacherID); err != nil {
		c.Flash().Add("danger", "action could not be completed")
		return c.Redirect(404, "/teacher/list")
	}
	if err := tx.Destroy(&teacher); err != nil {
		return err
	}
	c.Flash().Add("success", "teacher deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/teacher/list")
}
