package storage

import (
	"context"
	"log/slog"
	"strings"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

// LoginUser gets user's data from the database to check for correct credentials
func (s *storage) LoginUser(ctx context.Context, email string) (models.LoginUserResponse, error) {
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		slog.Error("Error while acquiring connection", slog.Any("error", err))
		return models.LoginUserResponse{}, err
	}
	defer conn.Release()

	var data models.LoginUserResponse

	query := `
		SELECT id, is_recruiter, password_hash, password_salt 
		FROM users 
		WHERE email = $1;
	`

	row := conn.QueryRow(ctx, query, strings.TrimSpace(email))

	if err := row.Scan(&data.UserID, &data.IsRecruiter, &data.PasswordHash, &data.PasswordSalt); err != nil {
		slog.Error("Error scanning", slog.Any("error", err))
		return models.LoginUserResponse{}, err
	}

	return data, nil
}
