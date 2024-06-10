package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

// Config provides service address and paths to the database.
type Config struct {
	HTTPAddress  string `env:"HTTP_ADDRESS" envDefault:":8080"`
	DatabasePath string `env:"DATABASE_URI"`
}

// New creates new Config
func New(logger *zap.Logger) *Config {
	var config = Config{}
	if err := env.Parse(&config); err != nil {
		logger.Error("Error occurred when parsing config", zap.Error(err))
	}

	flag.StringVar(&config.HTTPAddress, "h", config.HTTPAddress, "http launch address")
	flag.StringVar(&config.DatabasePath, "d", config.DatabasePath, "Path to database")
	flag.Parse()

	return &config
}
