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
	)

	api := e.Group("/api")

	// User group.
	userPath := api.Group("/user")
	userPath.POST("/register", nil)
	userPath.POST("/login", nil)

	return e
}
