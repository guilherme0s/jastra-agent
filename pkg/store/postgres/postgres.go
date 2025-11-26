package postgres

import (
	"fmt"
	"log/slog"

	"github.com/guilherme0s/atlans/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	DB *sqlx.DB
}

func New(cfg config.DatabaseSettings) (*Store, error) {
	cfg.SetDefaults()
	dsn := cfg.DSN()

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxConns)
	db.SetMaxIdleConns(cfg.MinConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	slog.Info("Successfully connected to the database")

	return &Store{DB: db}, nil
}

func (s *Store) Close() error {
	return s.DB.Close()
}
