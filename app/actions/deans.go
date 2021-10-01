package actions

import (
	"fmt"
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func NewDean(c buffalo.Context) error {
	c.Set("dean", models.Dean{})
	return c.Render(http.StatusOK, r.HTML("dean/new.plush.html"))
}

func CreateDean(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	dean := &models.Dean{}
	dean.Rol = "dean"
	if err := c.Bind(dean); err != nil {
		return err
	}
	verrs := dean.Validate(tx)
	if verrs.HasAny() {
		c.Set("dean", dean)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("dean/new.plush.html"))
	}
	if err := tx.Create(dean); err != nil {
		fmt.Println(err)
		return err
	}
	c.Flash().Add("success", "registered successfully")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListDeans(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	deans := models.Deans{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage
	if err := q.Order("first_name, last_name").All(&deans); err != nil {
		return err
	}
	c.Set("deans", deans)
	c.Set("paginatorD", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("dean/list.plush.html"))
}

func DeleteDean(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	dean := models.Dean{}
	deanID := c.Param("dean_id")
	if err := tx.Find(&dean, deanID); err != nil {
		c.Flash().Add("danger", "action could not be completed")
		return c.Redirect(404, "/dean/list")
	}
	if err := tx.Destroy(&dean); err != nil {
		return err
	}
	c.Flash().Add("success", "dean deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/dean/list")
}
