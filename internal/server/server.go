package server

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/SigmaMaleGroup/networkator/internal/config"
)

type Handlers interface {
	RegisterUser(c echo.Context) error
	LoginUser(c echo.Context) error
	CreateVacancy(c echo.Context) error
}

type Middleware interface {
	CheckToken(next echo.HandlerFunc) echo.HandlerFunc
	RequestLogger(next echo.HandlerFunc) echo.HandlerFunc
}

// server provides a single configuration out of all components
type server struct {
	httpHandlers Handlers
	middleware   Middleware
	config       *config.Config
}

// NewByConfig returns server instance with default config
func NewByConfig(
	httpHandlers Handlers,
	middleware Middleware,
	config *config.Config,
) *server {
	return &server{
		httpHandlers: httpHandlers,
		middleware:   middleware,
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
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := httpServ.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal("Error shutting down", err)
		}

		slog.Info("Server shut down", slog.String("http", s.config.HTTPAddress))
		serverStopCtx()
	}()

	slog.Info("Server started", slog.String("http", s.config.HTTPAddress))
	if err := httpServ.ListenAndServe(); err != nil {
		log.Fatal("Cant start server", err)
	}

	<-serverCtx.Done()
}
