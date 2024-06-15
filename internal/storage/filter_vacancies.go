package storage

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) GetVacanciesByFilter(
	ctx context.Context,
	filter models.VacancyFilterRequest,
) ([]models.VacancyShortInfo, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return nil, fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	queryFilters := make([]string, 0, 2)
	args := make([]any, 0, 2)
	query := `
      		SELECT name, salary_from, salary_to, city, employment_type, description 
      		FROM vacancies
      		WHERE
	`

	if filter.CompanyName != "" {
		args = append(args, filter.CompanyName)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s = $%d`, "company_name", len(args)))
	}

	if filter.City != "" {
		args = append(args, filter.City)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s = $%d`, "city", len(args)))
	}

	if filter.Experience > 0 {
		args = append(args, filter.Experience)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s = $%d`, "experience", len(args)))
	}

	if filter.EmploymentType > 0 {
		args = append(args, filter.EmploymentType)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s = $%d`, "employment_type", len(args)))
	}

	if filter.SalaryFrom > 0 {
		args = append(args, filter.SalaryFrom)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s > $%d`, "salary_from", len(args)))
	}

	if filter.SalaryTo > 0 {
		args = append(args, filter.SalaryTo)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s < $%d`, "salary_to", len(args)))
	}

	query += strings.Join(queryFilters, " AND ") + ";"

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	result := make([]models.VacancyShortInfo, 0, 4)

	for rows.Next() {
		var vacancy models.VacancyShortInfo

		if err := rows.Scan(
			&vacancy.Name,
			&vacancy.SalaryFrom,
			&vacancy.SalaryTo,
			&vacancy.City,
			&vacancy.EmploymentType,
			&vacancy.Description,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, vacancy)
	}

	return result, nil
}
