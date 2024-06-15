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
	vacancyPath.POST("/new", s.httpHandlers.CreateVacancy)
	vacancyPath.POST("/filter", s.httpHandlers.GetVacanciesByFilter)

	return e
}
