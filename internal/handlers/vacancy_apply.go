package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) VacancyApply(c echo.Context) error {
	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)
	userID, err := strconv.Atoi(userIDValue)
	if err != nil {
		slog.Error("Bad user id", slog.Any("error", err), slog.String("user_id", userIDValue))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if c.Request().Context().Value(models.CtxRoleKey).(string) != models.Applicant {
		slog.Error("user id not allowed", slog.String("user_id", userIDValue))
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	vacancyIDFromPath := c.Param("vacancyID")
	vacancyID, err := strconv.Atoi(vacancyIDFromPath)
	if err != nil {
		slog.Error("Bad vacancy id", slog.Any("error", err), slog.String("vacancy_id", vacancyIDFromPath))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := h.service.VacancyApply(c.Request().Context(), int64(vacancyID), int64(userID)); err != nil {
		slog.Error("Error in call to create vacancy", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")

	return nil
}
