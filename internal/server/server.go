package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/SigmaMaleGroup/networkator/internal/config"
)

type Handlers interface {
}

// server provides a single configuration out of all components
type server struct {
	httpHandlers Handlers
	config       *config.Config
	logger       *zap.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig(httpHandlers Handlers, logger *zap.Logger, config *config.Config) *server {
	return &server{
		httpHandlers: httpHandlers,
		logger:       logger,
		config:       config,
	}
}

// Run runs the service and provides graceful shutdown
func (s server) Run() {
	httpServ := &http.Server{Addr: s.config.HTTPAddress, Handler: s.Router()}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				s.logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := httpServ.Shutdown(shutdownCtx)
		if err != nil {
			s.logger.Fatal("Error shutting down", zap.Error(err))
		}

		s.logger.Info("Server shut down", zap.String("http", s.config.HTTPAddress))
		serverStopCtx()
	}()

	s.logger.Info("Server started", zap.String("http", s.config.HTTPAddress))
	if err := httpServ.ListenAndServe(); err != nil {
		s.logger.Fatal("Cant start server", zap.Error(err))
	}

	<-serverCtx.Done()
}
