package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ListTeacherFacultades(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	teachersFacultades := models.TeachersFacultades{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	status := c.Param("facultad")
	if status != "" {
		q.Where("facultad = ?", status)
	}

	if err := q.Order("facultad, nombre, apellido").All(&teachersFacultades); err != nil {
		return err
	}

	c.Set("teachersFacultades", teachersFacultades)
	c.Set("paginatorTF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("teacher/listTF.plush.html"))
}
