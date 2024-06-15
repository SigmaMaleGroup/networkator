package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) EditVacancy(ctx context.Context, vacancyID int64, vacancy models.VacancyRequest) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `
		UPDATE vacancies SET 
			recruiter_id = $1, 
			experience = $2, 
			city = $3, 
			employment_type = $4, 
			salary_from = $5, 
			salary_to = $6, 
			company_name = $7,
			name = $8,
			skills = $9,
			address = $10,
			description = $11        
		WHERE id = $12;
	`

	if _, err := conn.Exec(
		ctx,
		query,
		vacancy.RecruiterID,
		vacancy.Experience,
		vacancy.City,
		vacancy.EmploymentType,
		vacancy.SalaryFrom,
		vacancy.SalaryTo,
		vacancy.CompanyName,
		vacancy.Name,
		vacancy.Skills,
		vacancy.Address,
		vacancy.Description,
		vacancyID,
	); err != nil {
		slog.Error("error editing vacancy", slog.Any("error", err))
		return fmt.Errorf("could not edit vacancy: %w", err)
	}

	return nil
}
