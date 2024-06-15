package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) ResumesGetByFilter(ctx context.Context, filter models.ResumeFilterRequest) ([]models.Resume, error) {
	return s.storage.ResumesGetByFilter(ctx, filter)
}
