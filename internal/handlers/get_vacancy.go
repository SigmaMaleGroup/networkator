package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	errs "github.com/SigmaMaleGroup/networkator/internal/custom_errors"
)

func (h *handlers) GetVacancyByID(c echo.Context) error {
	vacancyIDFromPath := c.Param("vacancyID")

	vacancyID, err := strconv.Atoi(vacancyIDFromPath)
	if err != nil {
		slog.Error("Bad vacancy id", slog.Any("error", err), slog.String("vacancy_id", vacancyIDFromPath))
		return c.JSON(http.StatusBadRequest, nil)
	}

	resp, err := h.service.GetVacancyByID(c.Request().Context(), int64(vacancyID))
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return c.JSON(http.StatusNotFound, nil)
		}

		slog.Error("Error in call to get vacancy", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return c.JSON(http.StatusOK, resp)
}
