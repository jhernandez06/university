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
	deans := models.Deans{}
	DeansList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("first_name, last_name").All(&deans); err != nil {
		return err
	}
	for _, d := range deans {
		oneDean := map[string]interface{}{
			d.FirstName + " " + d.LastName: uuid.FromStringOrNil(d.ID.String()),
		}
		DeansList = append(DeansList, oneDean)
	}
	c.Set("deansList", DeansList)
	c.Set("faculty", models.Faculty{})
	return c.Render(http.StatusOK, r.HTML("faculty/new.plush.html"))
}

func CreateFaculty(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculty := models.Faculty{}
	deans := models.Deans{}
	DeansList := []map[string]interface{}{}

	q := tx.Q()
	if err := q.Order("first_name, last_name").All(&deans); err != nil {
		return err
	}
	for _, d := range deans {
		oneDean := map[string]interface{}{
			d.FirstName + " " + d.LastName: uuid.FromStringOrNil(d.ID.String()),
		}
		DeansList = append(DeansList, oneDean)
	}

	if err := c.Bind(&faculty); err != nil {
		return err
	}

	verrs := faculty.Validate(tx)
	if verrs.HasAny() {
		c.Set("faculty", faculty)
		c.Set("deansList", DeansList)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("faculty/new.plush.html"))
	}

	if err := tx.Create(&faculty); err != nil {
		return err
	}
	c.Flash().Add("success", "faculty create succesfully")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListFaculties(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculties := models.Faculties{}
	q := tx.PaginateFromParams(c.Params())
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := q.Order("name").All(&faculties); err != nil {
		return err
	}

	c.Set("faculties", faculties)
	c.Set("paginatorF", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("faculty/list.plush.html"))
}

func DeleteFaculty(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	faculty := models.Faculty{}
	facultyID := c.Param("faculty_id")
	if err := tx.Find(&faculty, facultyID); err != nil {
		c.Flash().Add("danger", "action could not be completed")
		return c.Redirect(404, "/faculty/list")
	}
	if err := tx.Destroy(&faculty); err != nil {
		return err
	}
	c.Flash().Add("success", "faculty deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/faculty/list")
}
