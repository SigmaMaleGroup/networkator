package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

// CheckToken implements JWT token parsing and authorizing users
func (m middleware) CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("auth")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			slog.Error("Error occurred while unpacking cookies", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		var signingKey = []byte(os.Getenv("secret"))

		if token != nil {
			token, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.Response().Status = http.StatusUnauthorized

					return "", err
				}

				return signingKey, nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID := claims["userID"].(int64)
				email := claims["email"].(string)
				role := claims["role"].(string)

				ctx := context.WithValue(c.Request().Context(), models.CtxUserIDKey, userID)
				ctx = context.WithValue(ctx, models.CtxEmailKey, email)
				ctx = context.WithValue(ctx, models.CtxRoleKey, role)

				c.SetRequest(c.Request().WithContext(ctx))
				return next(c)
			}
		}

		return c.JSON(http.StatusUnauthorized, nil)
	}
}
