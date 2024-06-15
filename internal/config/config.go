package config

import (
	"flag"
	"log/slog"

	"github.com/caarlos0/env/v6"
)

// Config provides service address and paths to the database.
type Config struct {
	HTTPAddress  string `env:"HTTP_ADDRESS" envDefault:":8080"`
	DatabasePath string `env:"DATABASE_URI"`
	Domain       string `env:"DOMAIN"`
}

// New creates new Config
func New() *Config {
	var config = Config{}
	if err := env.Parse(&config); err != nil {
		slog.Error("Error occurred when parsing config", err)
	}

	flag.StringVar(&config.HTTPAddress, "h", config.HTTPAddress, "http launch address")
	flag.StringVar(&config.DatabasePath, "s", config.DatabasePath, "Path to database")
	flag.StringVar(&config.DatabasePath, "d", config.Domain, "Domain")
	flag.Parse()

	return &config
}
