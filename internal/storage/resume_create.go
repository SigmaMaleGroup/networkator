package storage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

func (s *storage) ResumeCreate(ctx context.Context, userID int64, resume models.Resume) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		slog.Error("Error occurred creating tx", slog.Any("error", err))
		return err
	}

	query := `
		INSERT INTO resume (
		                    user_id, 
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
		                    ) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id;
	`

	row := tx.QueryRow(
		ctx,
		query,
		userID,
		resume.Fio,
		resume.Position,
		resume.Gender,
		resume.Address,
		resume.BirthDate,
		resume.Phone,
		resume.SalaryFrom,
		resume.SalaryTo,
		resume.Education,
		resume.Skills,
		resume.Nationality,
		resume.Disabilities,
	)
	var resumeID int64

	if err := row.Scan(&resumeID); err != nil {
		slog.Error("Error scanning", slog.Any("error", err))
		return errors.Join(err, tx.Rollback(ctx))
	}

	experienceQuery := `
		INSERT INTO experience (resume_id, company_name, time_from, time_to, position, work_exp_description) 
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	for _, val := range resume.WorkExperience {
		if _, err := tx.Exec(
			ctx,
			experienceQuery,
			resumeID,
			val.CompanyName,
			val.TimeFrom,
			val.TimeTo,
			val.Position,
			val.Description,
		); err != nil {
			slog.Error("Error exec experience", slog.Any("error", err))
			return errors.Join(err, tx.Rollback(ctx))
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		slog.Error("Error occurred making reservation in db", slog.Any("error", err))
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
