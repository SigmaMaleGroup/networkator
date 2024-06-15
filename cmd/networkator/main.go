package main

import (
	"log/slog"
	"os"

	"github.com/SigmaMaleGroup/networkator/internal/config"
	"github.com/SigmaMaleGroup/networkator/internal/handlers"
	"github.com/SigmaMaleGroup/networkator/internal/middleware"

	"github.com/SigmaMaleGroup/networkator/internal/server"
	"github.com/SigmaMaleGroup/networkator/internal/service"
	"github.com/SigmaMaleGroup/networkator/internal/storage"
)

func main() {
	slog.SetDefault(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	)

	cfg := config.New()

	store := storage.New(cfg.DatabasePath)
	store.CreateSchema()

	serv := service.New(store)
	httpHandle := handlers.New(serv, cfg.Domain)

	mdware := middleware.New()

	srv := server.NewByConfig(httpHandle, mdware, cfg)
	srv.Run()
}
