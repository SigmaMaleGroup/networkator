package storage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	errs "github.com/SigmaMaleGroup/networkator/internal/custom_errors"
	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) GetVacancyByID(ctx context.Context, vacancyID int64) (models.VacancyFullInfo, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return models.VacancyFullInfo{}, fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `
      		SELECT 
      		    id, 
      		    name, 
      		    city, 
      		    salary_from, 
      		    salary_to, 
      		    skills,
      		    experience,
      		    address,
      		    description,
      		    employment_type
      		FROM vacancies
      		WHERE id = $1;
	`
	var vacancy models.VacancyFullInfo

	row := conn.QueryRow(ctx, query, vacancyID)

	if err := row.Scan(
		&vacancy.ID,
		&vacancy.Name,
		&vacancy.City,
		&vacancy.SalaryFrom,
		&vacancy.SalaryTo,
		&vacancy.Skills,
		&vacancy.Experience,
		&vacancy.Address,
		&vacancy.Description,
		&vacancy.EmploymentType,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.VacancyFullInfo{}, fmt.Errorf("not found: %w", errs.ErrNotFound)
		}

		return models.VacancyFullInfo{}, fmt.Errorf("scan row: %w", err)
	}

	return vacancy, nil
}
