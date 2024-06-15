package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	errs "github.com/SigmaMaleGroup/networkator/internal/custom_errors"
	"github.com/SigmaMaleGroup/networkator/internal/models"
)

// LoginUser handles user login operations
func (h *handlers) LoginUser(c echo.Context) error {
	var req models.LoginCredentials

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if req.Email == "" || req.Password == "" {
		slog.Info("Credentials empty")
		return c.JSON(http.StatusBadRequest, nil)
	}

	token, err := h.service.LoginUser(c.Request().Context(), req.Email, req.Password)

	switch {
	case errors.Is(err, errs.ErrWrongCredentials):
		return c.JSON(http.StatusBadRequest, nil)
	case err != nil:
		slog.Error("Error in call to processor", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")

	c.SetCookie(&http.Cookie{
		Name:   "auth",
		Value:  token,
		Secure: false,
		Domain: h.domain,
		Path:   "/",
	})

	return nil
}
