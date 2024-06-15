package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) GetVacancyByID(ctx context.Context, vacancyID int64) (models.VacancyFullInfo, error) {
	return s.storage.GetVacancyByID(ctx, vacancyID)
}
