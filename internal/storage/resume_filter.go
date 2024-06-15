package storage

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) ResumesGetByFilter(ctx context.Context, filter models.ResumeFilterRequest) ([]models.Resume, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return nil, fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	queryFilters := make([]string, 0, 2)
	args := make([]any, 0, 2)
	query := `
      		SELECT
			id,
			fio, 
			job_name, 
			gender, 
			address, 
			birth_date, 
			phone_number, 
			salary_from, 
			salary_to, 
			education, 
			skills, 
			nationality, 
			disability
		FROM resume 
		WHERE 
	`

	if filter.Education {
		queryFilters = append(queryFilters, fmt.Sprintf(`%s != ''`, "education"))
	}

	if filter.SalaryFrom > 0 {
		args = append(args, filter.SalaryFrom)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s >= $%d`, "salary_from", len(args)))
	}

	if filter.SalaryTo > 0 {
		args = append(args, filter.SalaryTo)
		queryFilters = append(queryFilters, fmt.Sprintf(`%s <= $%d`, "salary_to", len(args)))
	}

	if len(filter.Skills) > 0 {
		quotedSlice := make([]string, len(filter.Skills))
		for i, str := range filter.Skills {
			quotedSlice[i] = fmt.Sprintf("'%s'", str)
		}
		result := strings.Join(quotedSlice, ",")

		queryFilters = append(queryFilters, fmt.Sprintf(`@> ARRAY[%s]`, result))
	}

	query += strings.Join(queryFilters, " AND ") + ";"

	if !filter.Education && filter.SalaryFrom == 0 && filter.SalaryTo == 0 && len(filter.Skills) == 0 {
		query = `
			SELECT
				id,
				fio, 
				job_name, 
				gender, 
				address, 
				birth_date, 
				phone_number, 
				salary_from, 
				salary_to, 
				education, 
				skills, 
				nationality, 
				disability
			FROM resume 
		`
	}

	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	result := make([]models.Resume, 0, 4)

	for rows.Next() {
		var resume models.Resume

		if err := rows.Scan(
			&resume.ID,
			&resume.Fio,
			&resume.Position,
			&resume.Gender,
			&resume.Address,
			&resume.BirthDate,
			&resume.Phone,
			&resume.SalaryFrom,
			&resume.SalaryTo,
			&resume.Education,
			&resume.Skills,
			&resume.Nationality,
			&resume.Disabilities,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		result = append(result, resume)
	}

	return result, nil
}
