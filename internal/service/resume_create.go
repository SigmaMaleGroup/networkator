package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) ResumeCreate(ctx context.Context, userID int64, resume models.Resume) error {
	return s.storage.ResumeCreate(ctx, userID, resume)
}
