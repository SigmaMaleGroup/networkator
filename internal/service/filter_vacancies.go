package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *service) GetVacanciesByFilter(
	ctx context.Context,
	filter models.VacancyFilterRequest,
) (models.VacancyFilterResponse, error) {
	if filter.EmploymentType == 0 && filter.Experience == 0 && filter.SalaryFrom == 0 &&
		filter.SalaryTo == 0 && filter.City == "" && filter.CompanyName == "" {
		return models.VacancyFilterResponse{}, errors.New("at least one filter must be filled")
	}

	res, err := s.storage.GetVacanciesByFilter(ctx, filter)
	if err != nil {
		return models.VacancyFilterResponse{}, fmt.Errorf("get vacancies: %w", err)
	}

	return models.VacancyFilterResponse{Vacancies: res}, nil
}
