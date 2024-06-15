package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router returns service's echo router
func (s server) Router() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.Gzip(),
		s.middleware.RequestLogger,
	)

	api := e.Group("/api")

	// User group.
	userPath := api.Group("/user")
	userPath.POST("/register", s.httpHandlers.RegisterUser)
	userPath.POST("/login", s.httpHandlers.LoginUser)

	// Vacancy group.
	vacancyPath := api.Group("/vacancy")
	vacancyPath.Use(s.middleware.CheckToken)

	vacancyPath.POST("/new", s.httpHandlers.CreateVacancy)
	vacancyPath.POST("/filter", s.httpHandlers.GetVacanciesByFilter)
	vacancyPath.GET("/:vacancyID", s.httpHandlers.GetVacancyByID)
	vacancyPath.POST("/edit/:vacancyID", s.httpHandlers.EditVacancy)
	vacancyPath.POST("/archive/:vacancyID", s.httpHandlers.ArchiveVacancy)
	vacancyPath.POST("/apply/:vacancyID", s.httpHandlers.VacancyApply)

	// Resume group.
	resumePath := api.Group("/resume")
	resumePath.Use(s.middleware.CheckToken)
	resumePath.POST("/new", s.httpHandlers.ResumeCreate)
	resumePath.POST("/all", nil)
	resumePath.POST("/:userID", nil)

	return e
}
