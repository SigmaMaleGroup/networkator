package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) GetVacanciesByFilter(c echo.Context) error {
	var req models.VacancyFilterRequest

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	resp, err := h.service.GetVacanciesByFilter(c.Request().Context(), req)
	if err != nil {
		slog.Error("Error in call to filter vacancies", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return c.JSON(http.StatusOK, resp)
}
