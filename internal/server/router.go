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
		//middleware.BasicAuth(),
	)

	return e
}
