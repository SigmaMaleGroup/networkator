package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) CreateVacancy(ctx context.Context, vacancy models.NewVacancyRequest) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `
		INSERT INTO vacancies (
		                       recruiter_id, 
		                       experience, 
		                       city, 
		                       employment_type, 
		                       salary_from, 
		                       salary_to, 
		                       company_name,
		                       name,
		                       skills,
		                       address,
		                       description
		                       ) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
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
	); err != nil {
		slog.Error("error inserting vacancy", slog.Any("error", err))
		return fmt.Errorf("could not insert vacancy: %w", err)
	}

	return nil
}
