package postgres

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/guilherme0s/atlans/pkg/config"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func (s *Store) MigrateUp(cfg config.DatabaseSettings) error {
	driver, err := postgres.WithInstance(s.DB.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver for migration: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.MigrationsPath,
		cfg.DBName, driver,
	)

	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not apply migrations: %w", err)
	}

	slog.Info("Database migrations applied successfully")

	return nil
}
