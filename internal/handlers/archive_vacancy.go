package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handlers) ArchiveVacancy(c echo.Context) error {
	vacancyIDFromPath := c.Param("vacancyID")

	vacancyID, err := strconv.Atoi(vacancyIDFromPath)
	if err != nil {
		slog.Error("Bad vacancy id", slog.Any("error", err), slog.String("vacancy_id", vacancyIDFromPath))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := h.service.ArchiveVacancy(c.Request().Context(), int64(vacancyID)); err != nil {
		slog.Error("Error in call to archive vacancy", slog.Any("error", err))
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return nil
}
