package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) StagesGet(
	ctx context.Context,
	request models.GetUsersStagesRequest,
) ([]models.ResumeForStages, error) {
	return s.storage.StagesGet(ctx, request)
}
