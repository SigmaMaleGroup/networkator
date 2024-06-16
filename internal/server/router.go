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
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"https://" + s.config.Domain, "http://" + s.config.Domain},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
			ExposeHeaders:    []string{"Link"},
			MaxAge:           300,
		}),
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
	resumePath.POST("/filter", s.httpHandlers.ResumesGetByFilter)
	resumePath.GET("/:userID", s.httpHandlers.ResumeGet)

	// Stages group.
	stagesPath := api.Group("/stages")
	stagesPath.Use(s.middleware.CheckToken)

	stagesPath.POST("/", s.httpHandlers.StagesGet)

	return e
}
