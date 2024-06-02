package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

var schema = `

`

// New creates a new instance of database layer using Postgres
func New(path string, logger *zap.Logger) *storage {
	// Wait until database initialize in container
	time.Sleep(time.Second * 2)

	config, err := pgxpool.ParseConfig(path)
	if err != nil {
		logger.Fatal("Unable to parse config", zap.Error(err))
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}

	return &storage{
		pool:   conn,
		logger: logger,
	}
}

// CreateSchema executes needed schema
func (s storage) CreateSchema() {
	_, err := s.pool.Exec(context.Background(), schema)
	if err != nil {
		s.logger.Fatal("Error occurred while executing schema", zap.Error(err))
	}

	s.logger.Info("Launched with Postgres")
}
