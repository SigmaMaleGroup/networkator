package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) ResumeGet(ctx context.Context, userID int64) (models.Resume, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return models.Resume{}, fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

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
		WHERE id = $1;
	`

	var (
		resumeID   int64
		resume     models.Resume
		experience = make([]models.Experience, 0, 4)
	)

	if err := conn.QueryRow(ctx, query, userID).Scan(
		&resumeID,
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
		slog.Error("Error scanning", slog.Any("error", err))
		return models.Resume{}, fmt.Errorf("scan resume: %w", err)
	}

	experienceQuery := `
		SELECT company_name, time_from, time_to, position, work_exp_description
		FROM experience
		WHERE resume_id = $1;
	`

	rows, err := conn.Query(ctx, experienceQuery, resumeID)
	if err != nil {
		slog.Error("Error exec experience", slog.Any("error", err))
		return models.Resume{}, fmt.Errorf("query experience: %w", err)
	}

	for rows.Next() {
		var exp models.Experience

		if err := rows.Scan(
			&exp.CompanyName,
			&exp.TimeFrom,
			&exp.TimeTo,
			&exp.Position,
			&exp.Description,
		); err != nil {
			slog.Error("Error scan experience", slog.Any("error", err))
			return models.Resume{}, fmt.Errorf("scan experience: %w", err)
		}

		experience = append(experience, exp)
	}
	resume.WorkExperience = experience

	return resume, nil
}
