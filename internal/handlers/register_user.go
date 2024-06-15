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

// RegisterUser handles user registration operations
func (h *handlers) RegisterUser(c echo.Context) error {
	var req models.RegisterCredentials

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if req.Password == "" || req.Email == "" {
		slog.Info("Credentials empty")
		return c.JSON(http.StatusBadRequest, nil)
	}

	token, err := h.service.RegisterUser(c.Request().Context(), &req)

	switch {
	case errors.Is(err, errs.ErrCredentialsInUse):
		return c.JSON(http.StatusConflict, nil)
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

	return c.JSON(http.StatusOK, models.AuthResponse{Token: token})
}
