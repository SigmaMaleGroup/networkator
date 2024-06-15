package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) ResumeGet(ctx context.Context, userID int64) (models.Resume, error) {
	return s.storage.ResumeGet(ctx, userID)
}
