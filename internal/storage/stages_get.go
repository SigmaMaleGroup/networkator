package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) StagesGet(
	ctx context.Context,
	request models.GetUsersStagesRequest,
) ([]models.ResumeForStages, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return nil, fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `
		SELECT
			r.id,
			r.user_id,
			r.fio, 
			r.job_name, 
			r.gender, 
			r.address, 
			r.birth_date, 
			r.phone_number, 
			r.salary_from, 
			r.salary_to, 
			r.education, 
			r.skills,
			r.nationality, 
			r.disability,
			a.stage
		FROM resume r
		JOIN applications a ON r.user_id = a.user_id
		WHERE a.vacancy_id = $1 AND a.stage = $2;
	`

	var resumes []models.ResumeForStages

	rows, err := conn.Query(ctx, query, request.VacancyID, request.StageName)
	if err != nil {
		slog.Error("Error query", slog.Any("error", err))
		return nil, fmt.Errorf("query: %w", err)
	}

	for rows.Next() {
		var resume models.ResumeForStages

		if err := rows.Scan(
			&resume.ID,
			&resume.UserID,
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
			&resume.StageName,
		); err != nil {
			slog.Error("Error scanning", slog.Any("error", err))
			return nil, fmt.Errorf("scan: %w", err)
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}
