package handlers

import (
	"context"

	"github.com/SigmaMaleGroup/networkator/internal/models"
)

type Service interface {
	LoginUser(ctx context.Context, email, password string) (string, error)
	RegisterUser(ctx context.Context, credits *models.RegisterCredentials) (string, error)
	CreateVacancy(ctx context.Context, vacancy models.NewVacancyRequest) error
	GetVacanciesByFilter(ctx context.Context, filter models.VacancyFilterRequest) (models.VacancyFilterResponse, error)
}

// handlers provides http-handlers for service
type handlers struct {
	service Service
	domain  string
}

// New creates new instance of handlers
func New(service Service, domain string) *handlers {
	return &handlers{
		service: service,
		domain:  domain,
	}
}
