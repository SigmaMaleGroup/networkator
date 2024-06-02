package service

import (
	"go.uber.org/zap"
)

type Storage interface {
}

// service provides business-logic
type service struct {
	storage Storage
	logger  *zap.Logger
}

// New creates new instance of actions
func New(storage Storage, logger *zap.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}
