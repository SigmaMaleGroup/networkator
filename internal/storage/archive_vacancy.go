package storage

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *storage) ArchiveVacancy(ctx context.Context, vacancyID int64) error {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return fmt.Errorf("could not acquire conn: %w", err)
	}
	defer conn.Release()

	query := `UPDATE vacancies SET archived = true WHERE id = $1;`

	if _, err := conn.Exec(ctx, query, vacancyID); err != nil {
		slog.Error("error inserting vacancy", slog.Any("error", err))
		return fmt.Errorf("could not insert vacancy: %w", err)
	}

	return nil
}
