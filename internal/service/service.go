package service

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

type Storage interface {
	CheckDuplicateUser(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email, passwordHash, passwordSalt string, isRecruiter bool) (int64, error)
	LoginUser(ctx context.Context, email string) (models.LoginUserResponse, error)
}

// service provides business-logic
type service struct {
	storage Storage
}

// New creates new instance of actions
func New(storage Storage) *service {
	return &service{
		storage: storage,
	}
}
