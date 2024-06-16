package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) StagesGet(c echo.Context) error {
	var req models.GetUsersStagesRequest

	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)
	if c.Request().Context().Value(models.CtxRoleKey).(string) != models.Recruiter {
		slog.Error("user id not allowed", slog.String("user_id", userIDValue))
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		slog.Error("Unable to read body", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if err = json.Unmarshal(body, &req); err != nil {
		slog.Error("Unable to decode JSON", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if req.StageName == "" || req.VacancyID == 0 {
		slog.Info("Empty request")
		return c.JSON(http.StatusBadRequest, nil)
	}

	resp, err := h.service.StagesGet(c.Request().Context(), req)
	if err != nil {
		slog.Error("Error in call to get stages", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return c.JSON(http.StatusOK, resp)
}
