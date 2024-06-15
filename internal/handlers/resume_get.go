package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (h *handlers) ResumeGet(c echo.Context) error {
	userIDValue := c.Request().Context().Value(models.CtxUserIDKey).(string)

	userID, err := strconv.Atoi(userIDValue)
	if err != nil || userID == 0 {
		slog.Error("Bad user id", slog.Any("error", err), slog.String("user_id", userIDValue))
		return c.JSON(http.StatusBadRequest, nil)
	}

	if c.Request().Context().Value(models.CtxRoleKey).(string) != models.Recruiter {
		slog.Error("user id not allowed", slog.String("user_id", userIDValue))
		return c.JSON(http.StatusMethodNotAllowed, nil)
	}

	applicantIDFromPath := c.Param("userID")

	applicantID, err := strconv.Atoi(applicantIDFromPath)
	if err != nil {
		slog.Error(
			"Bad applicant id",
			slog.Any("error", err),
			slog.String("applicant_id", applicantIDFromPath),
		)
		return c.JSON(http.StatusBadRequest, nil)
	}

	resume, err := h.service.ResumeGet(c.Request().Context(), int64(applicantID))
	if err != nil {
		slog.Error("Error in call to create resume", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	c.Response().Header().Add("Access-Control-Allow-Credentials", "true")

	return c.JSON(http.StatusOK, resume)
}
