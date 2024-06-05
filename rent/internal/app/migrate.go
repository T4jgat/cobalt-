package app

import (
	"database/sql"
	"fmt"
	"github.com/T4jgat/cobalt+/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func RunMigrations(postgresDSN string, log logger.Logger) error {
	log.Info("Starting database migrations...")

	db, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://C:/Home/code/golangprojects/cobalt+/rent/migrations", // path to migration files
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %s", err)
	}

	log.Info("Database migrations completed successfully.")
	return nil
}
