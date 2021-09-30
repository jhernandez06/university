package actions

import (
	"fmt"
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func NewDecano(c buffalo.Context) error {
	c.Set("decano", models.Decano{})
	return c.Render(http.StatusOK, r.HTML("decano/new.plush.html"))
}

func CreateDecano(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	decano := &models.Decano{}
	decano.Rol = "decano"
	if err := c.Bind(decano); err != nil {
		return err
	}
	verrs := decano.Validate(tx)
	if verrs.HasAny() {
		c.Set("decano", decano)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("decano/new.plush.html"))
	}
	if err := tx.Create(decano); err != nil {
		fmt.Println(err)
		return err
	}
	c.Flash().Add("success", "registered successfully")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListDecanos(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	decanos := models.Decanoes{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage
	if err := q.Order("nombre, apellido").All(&decanos); err != nil {
		return err
	}
	c.Set("decanos", decanos)
	c.Set("paginatorD", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("decano/list.plush.html"))
}
