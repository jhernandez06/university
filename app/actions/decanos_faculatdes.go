package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ListDecanosFacultades(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	decanosFacultades := models.DecanosFacultades{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("nombre").All(&decanosFacultades); err != nil {
		return err
	}

	c.Set("decanoFacultades", decanosFacultades)
	c.Set("paginatorDF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("facultad/listDF.plush.html"))
}
