package actions

import (
	"fmt"
	"net/http"
	"university/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

func NewCourse(c buffalo.Context) error {
	c.Set("course", models.Course{})
	return c.Render(http.StatusOK, r.HTML("course/new.plush.html"))
}

func CreateCourse(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	course := models.Course{}

	if err := c.Bind(&course); err != nil {
		return err
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
	q.Paginator.PerPage = 5
	q.Paginator.Offset = (q.Paginator.Page * q.Paginator.PerPage) - q.Paginator.PerPage

	if err := c.Bind(&course); err != nil {
		return err
	}
	if course.KeyWord != "" {
		key := fmt.Sprintf("name LIKE '%%%s%%'", course.KeyWord)

		if err := q.Where(key).All(&courses); err != nil {
			return err
		}
		count, err := q.Count(&courses)
		if err != nil {
			return err
		}

		c.Set("search", true)
		c.Set("count", count)
		c.Set("courses", courses)
		c.Set("course", course)
		c.Set("paginatorC", q.Paginator)
		return c.Render(http.StatusOK, r.HTML("course/list.plush.html"))
	}

	if err := q.Order("created_at,name").All(&courses); err != nil {
		return err
	}
	c.Set("search", false)
	c.Set("courses", courses)
	c.Set("course", course)
	c.Set("paginatorC", q.Paginator)
	return c.Render(http.StatusOK, r.HTML("course/list.plush.html"))
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
