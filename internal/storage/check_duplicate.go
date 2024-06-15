package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

// CheckDuplicateUser checks if user is already existing
func (s *storage) CheckDuplicateUser(ctx context.Context, email string) (bool, error) {
	var dbEmail string

	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return false, err
	}
	defer conn.Release()

	if err := conn.QueryRow(ctx, "SELECT email FROM users WHERE email = $1", email).Scan(&dbEmail); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("query: %w", err)
	}

	if dbEmail == email {
		return true, nil
	}

	return false, nil
}
