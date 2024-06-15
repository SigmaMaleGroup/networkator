package service

import (
	"context"

	errs "github.com/SigmaMaleGroup/networkator/internal/customErrors"
	"github.com/SigmaMaleGroup/networkator/internal/models"
	"github.com/SigmaMaleGroup/networkator/internal/tokens"
)

func (s service) LoginUser(ctx context.Context, email, password string) (string, error) {
	res, err := s.storage.LoginUser(ctx, email)
	if err != nil {
		return "", err
	}

	checkHash := mdHash(password, res.PasswordSalt)
	if checkHash != res.PasswordHash {
		return "", errs.ErrWrongCredentials
	}

	return tokens.GenerateJWT(res.UserID, email, models.GetUserType(res.IsRecruiter))
}
