package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) EditVacancy(ctx context.Context, vacancyID int64, vacancy models.VacancyRequest) error {
	return s.storage.EditVacancy(ctx, vacancyID, vacancy)
}
