package main

import (
	"github.com/SigmaMaleGroup/networkator/internal/config"
	"github.com/SigmaMaleGroup/networkator/internal/handlers"
	"github.com/SigmaMaleGroup/networkator/internal/logger"
	"github.com/SigmaMaleGroup/networkator/internal/server"
	"github.com/SigmaMaleGroup/networkator/internal/service"
	"github.com/SigmaMaleGroup/networkator/internal/storage"
)

func main() {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	store := storage.New(cfg.DatabasePath, log)
	//store.CreateSchema()

	serv := service.New(store, log)
	httpHandle := handlers.New(serv, log)

	srv := server.NewByConfig(httpHandle, log, cfg)
	srv.Run()
}
