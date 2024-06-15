package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func (m middleware) RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		slog.Debug("Request",
			slog.Any("method", c.Request().Method),
			slog.Any("host", c.Request().Host),
			slog.Any("remote addr", c.Request().RemoteAddr),
			slog.Any("req uri", c.Request().RequestURI),
			slog.Any("header", c.Request().Header),
		)

		return next(c)
	}
}
