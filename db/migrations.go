package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const migrationVersion = 1

//go:embed migrations/*.sql
var migrations embed.FS

func buildMigrationClient(db *sql.DB) (*migrate.Migrate, error) {
	source, err := iofs.New(migrations, "migrations")

	if err != nil {
		return nil, fmt.Errorf("failed to read migration source %w", err)
	}

	target, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return nil, fmt.Errorf("failed to read migration target %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", source, "postgres", target)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance %w", err)
	}

	return m, nil
}

func RunUpMigrations(db *sql.DB, logger *log.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Fatal("failed to build client for up migration", err)
		return err
	}

	err = m.Migrate(migrationVersion)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal("failed to run up migration", err, fmt.Sprintf("version: %s", migrationVersion))
		return err
	}

	defer func(m *migrate.Migrate) {
		sErr, tErr := m.Close()
		if sErr != nil {
			logger.Fatal("Failed to close source", err)
		}

		if tErr != nil {
			logger.Fatal("Failed to close target", err)
		}
	}(m)

	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Database is already migrated. Nothing to change", fmt.Sprintf("version: %s", migrationVersion))
		return nil
	}

	fmt.Println("Successfully executed migrations for DB", fmt.Sprintf("version: %s", migrationVersion))
	return err
}

func RunDownMigration(db *sql.DB, logger *log.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Fatal("failed to build client for down migration", err)
		return err
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal("failed to run migration", err, fmt.Sprint("version: %s", migrationVersion))
		return err
	}

	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Database is already down migrated. Nothing to change", fmt.Sprintf("version: %s", migrationVersion))
		return nil
	}

	fmt.Println("Database is now reset")

	return nil
}
