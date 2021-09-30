package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ListTeacherCourses(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	teacherCourses := models.ViewTeacherCourses{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("facultad, nombre, apellido").All(&teacherCourses); err != nil {
		return err
	}

	c.Set("teacherCourses", teacherCourses)
	c.Set("paginatorTC", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("teacher/listTC.plush.html"))
}
