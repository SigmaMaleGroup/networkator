package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) CreateVacancy(ctx context.Context, vacancy models.VacancyRequest) error {
	return s.storage.CreateVacancy(ctx, vacancy)
}
