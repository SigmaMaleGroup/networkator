package handlers

import (
	"go.uber.org/zap"
)

type Service interface {
}

// handlers provides http-handlers for service
type handlers struct {
	service Service
	logger  *zap.Logger
}

// New creates new instance of handlers
func New(service Service, logger *zap.Logger) *handlers {
	return &handlers{
		service: service,
		logger:  logger,
	}
}
