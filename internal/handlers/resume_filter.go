package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) ResumesGetByFilter(c echo.Context) error {
	var req models.ResumeFilterRequest

	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)
	if c.Request().Context().Value(models.CtxRoleKey).(string) != models.Recruiter {
		slog.Error("user id not allowed", slog.String("user_id", userIDValue))
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	resp, err := h.service.ResumesGetByFilter(c.Request().Context(), req)
	if err != nil {
		slog.Error("Error in call to filter resumes", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, resp)
}
