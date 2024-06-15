package storage

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *storage) VacancyApply(ctx context.Context, vacancyID, userID int64) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `
		INSERT INTO applications (vacancy_id, user_id) 
		VALUES ($1, $2);
	`

	if _, err := conn.Exec(ctx, query, vacancyID, userID); err != nil {
		slog.Error("error apply vacancy", slog.Any("error", err))
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
