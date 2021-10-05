package actions

import (
	"fmt"
	"net/http"
	"strconv"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
)

func NewCourse(c buffalo.Context) error {
	c.Set("course", models.Course{})
	return c.Render(http.StatusOK, r.HTML("course/new.plush.html"))
}

func CreateCourse(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course := models.Course{}

	if err := c.Bind(&course); err != nil {
		verrs := validate.NewErrors()
		verrs.Add("creditos", "SE JODIO LOLA XD")
		c.Set("course", course)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("course/new.plush.html"))
	}

	verrs := course.Validate(tx)
	if verrs.HasAny() {
		c.Set("course", course)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("course/new.plush.html"))
	}

	if err := tx.Create(&course); err != nil {
		return err
	}

	c.Flash().Add("success", "Successfully created course")
	return c.Redirect(http.StatusSeeOther, "/")
}

func ListCourses(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	courses := models.Courses{}
	course := models.Course{}
	q := tx.PaginateFromParams(c.Params())
	orderValues := []string{"created_at", "name", "name desc", "cast(code as int) ASC", "cast(code as int) DESC"}
	order := 0
	if paramOrderBy, err := strconv.Atoi(c.Param("order")); err == nil {
		if paramOrderBy >= len(orderValues) || paramOrderBy < 0 {
			fmt.Println(len(orderValues))
		}
		order = paramOrderBy

	}
	perPage := 5
	if paramPerPage, err := strconv.Atoi(c.Param("perPage")); err == nil {
		perPage = paramPerPage
	}
	q.Paginator.PerPage = perPage
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := c.Bind(&course); err != nil {
		return err
	}

	List := func(c buffalo.Context) error {
		c.Set("courses", courses)
		c.Set("course", course)
		c.Set("perPage", q.Paginator.PerPage)
		c.Set("paginatorC", q.Paginator)
		return c.Render(http.StatusOK, r.HTML("course/list.plush.html"))
	}

	if course.KeyWord != "" {
		key := "name LIKE ? OR "
		key += "code LIKE ?"
		if err := q.Where(key, "%"+course.KeyWord+"%", "%"+course.KeyWord+"%").All(&courses); err != nil {
			return err
		}
		count, err := q.Count(&courses)
		if err != nil {
			return err
		}

		c.Set("count", count)
		c.Set("search", true)
		return List(c)
	}

	if err := q.Order(orderValues[order]).All(&courses); err != nil {
		return err
	}
	fmt.Println("------------------------------>>>>", c.Request().URL)
	c.Set("search", false)
	return List(c)
}

func DeleteCourse(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course := models.Course{}
	courseID := c.Param("course_id")
	if err := tx.Find(&course, courseID); err != nil {
		c.Flash().Add("danger", "action could not be completed")
		return c.Redirect(404, "/course/list")
	}
	if err := tx.Destroy(&course); err != nil {
		return err
	}
	c.Flash().Add("success", "course deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/course/list")
}
