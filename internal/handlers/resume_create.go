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

func (h *handlers) ResumeCreate(c echo.Context) error {
	var req models.Resume

	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)

	userID, err := strconv.Atoi(userIDValue)
	if err != nil || userID == 0 {
		slog.Error("Bad user id", slog.Any("error", err), slog.String("user_id", userIDValue))
		return c.JSON(http.StatusBadRequest, nil)
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

	if err := h.service.ResumeCreate(c.Request().Context(), int64(userID), req); err != nil {
		slog.Error("Error in call to create resume", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")
	c.Response().Header().Add("Access-Control-Allow-Origin", "https://"+h.domain)

	return nil
}
