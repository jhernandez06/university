package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func ListDeansFaculties(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	deanFaculties := models.DeanFaculties{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("first_name").All(&deanFaculties); err != nil {
		return err
	}

	c.Set("deanFaculties", deanFaculties)
	c.Set("paginatorDF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("faculty/listDF.plush.html"))
}
