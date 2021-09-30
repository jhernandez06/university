package actions

import (
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func NewFaculty(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	decanos := models.Decanoes{}
	DecanosList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("nombre, apellido").All(&decanos); err != nil {
		return err
	}
	for _, decano := range decanos {
		oneDecano := map[string]interface{}{
			decano.Nombre + " " + decano.Apellido: uuid.FromStringOrNil(decano.ID.String()),
		}
		DecanosList = append(DecanosList, oneDecano)
	}
	c.Set("decanosList", DecanosList)
	c.Set("faculty", models.Facultad{})
	return c.Render(http.StatusOK, r.HTML("facultad/new.plush.html"))
}

func CreateFaculty(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculty := models.Facultad{}
	decanos := models.Decanoes{}
	DecanosList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("nombre, apellido").All(&decanos); err != nil {
		return err
	}
	for _, decano := range decanos {
		oneDecano := map[string]interface{}{
			decano.Nombre + " " + decano.Apellido: uuid.FromStringOrNil(decano.ID.String()),
		}
		DecanosList = append(DecanosList, oneDecano)
	}

	if err := c.Bind(&faculty); err != nil {
		return err
	}

	verrs := faculty.Validate(tx)
	if verrs.HasAny() {
		c.Set("faculty", faculty)
		c.Set("decanosList", DecanosList)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("facultad/new.plush.html"))
	}

	if err := tx.Create(&faculty); err != nil {
		return err
	}
	c.Flash().Add("success", "faculty create succesfully")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListFaculties(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculties := models.Facultades{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("nombre").All(&faculties); err != nil {
		return err
	}

	c.Set("faculties", faculties)
	c.Set("paginatorF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("facultad/list.plush.html"))
}
