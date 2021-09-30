package app

import (
	base "university"
	"university/app/actions"
	"university/app/actions/home"
	"university/app/middleware"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)

	root.GET("/decano/new", actions.NewDecano)
	root.POST("/decano/create", actions.CreateDecano)
	root.GET("/decano/list", actions.ListDecanos)

	root.GET("/faculty/new", actions.NewFaculty)
	root.POST("/faculty/create", actions.CreateFaculty)
	root.GET("/faculty/list", actions.ListFaculties)
	root.GET("/faculty/listDF", actions.ListDecanosFacultades)

	root.GET("/teacher/new", actions.NewTeacher)
	root.POST("/teacher/create", actions.CreateTeacher)
	root.GET("/teacher/list", actions.ListTeacher)
	root.GET("/teacher/listTF", actions.ListTeacherFacultades)
	root.GET("/teacher/listTC", actions.ListTeacherCourses)

	root.GET("/course/new", actions.NewCourse)
	root.POST("/course/create", actions.CreateCourse)
	root.GET("/course/list", actions.ListCourses)

	root.ServeFiles("/", base.Assets)
}
