package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) CreateVacancy(c echo.Context) error {
	var req models.VacancyRequest
	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)

	userID, err := strconv.Atoi(userIDValue)
	if err != nil {
		slog.Error("Bad user id", slog.Any("error", err), slog.String("user_id", userIDValue))
		return c.JSON(http.StatusBadRequest, nil)
	}
	req.RecruiterID = int64(userID)

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err := h.service.CreateVacancy(c.Request().Context(), req); err != nil {
		slog.Error("Error in call to create vacancy", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return nil
}
