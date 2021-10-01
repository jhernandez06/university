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

	root.GET("/dean/new", actions.NewDean)
	root.POST("/dean/create", actions.CreateDean)
	root.GET("/dean/list", actions.ListDeans)
	root.DELETE("/dean/delete/{dean_id}", actions.DeleteDean)

	root.GET("/faculty/new", actions.NewFaculty)
	root.POST("/faculty/create", actions.CreateFaculty)
	root.GET("/faculty/list", actions.ListFaculties)
	root.GET("/faculty/listDF", actions.ListDeansFaculties)
	root.DELETE("/faculty/delete/{faculty_id}", actions.DeleteFaculty)

	root.GET("/teacher/new", actions.NewTeacher)
	root.POST("/teacher/create", actions.CreateTeacher)
	root.GET("/teacher/list", actions.ListTeacher)
	root.GET("/teacher/listTF", actions.ListTeacherFaculties)
	root.GET("/teacher/listTC", actions.ListTeacherCourses)
	root.DELETE("/teacher/delete/{teacher_id}", actions.DeleteTeacher)

	root.GET("/course/new", actions.NewCourse)
	root.POST("/course/create", actions.CreateCourse)
	root.GET("/course/list", actions.ListCourses)
	root.DELETE("/course/delete/{course_id}", actions.DeleteCourse)

	root.ServeFiles("/", base.Assets)
}
