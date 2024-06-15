package storage

import (
	"context"
	"log/slog"
	"strings"
)

// CreateUser inserts new user's data in the database
func (s *storage) CreateUser(ctx context.Context, email, passwordHash, passwordSalt string, isRecruiter bool) (int64, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", err)
		return 0, err
	}
	defer conn.Release()

	query := `
		INSERT INTO users (email, password_hash, password_salt, is_recruiter) 
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	row := conn.QueryRow(ctx, query, strings.TrimSpace(email), passwordHash, passwordSalt, isRecruiter)
	var userID int64

	if err := row.Scan(&userID); err != nil {
		slog.Error("Error scanning", err)
		return 0, err
	}

	return userID, nil
}
