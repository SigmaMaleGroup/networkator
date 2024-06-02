package logger

import (
	"go.uber.org/zap"
)

// InitializeLogger initializes new Logger configured instance
func InitializeLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
