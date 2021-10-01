package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ListTeacherFaculties(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	teacherFaculties := models.TeacherFaculties{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	status := c.Param("faculty")
	if status != "" {
		q.Where("faculty = ?", status)
	}

	if err := q.Order("faculty, first_name, last_name").All(&teacherFaculties); err != nil {
		return err
	}

	c.Set("teacherFaculties", teacherFaculties)
	c.Set("paginatorTF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("teacher/listTF.plush.html"))
}
