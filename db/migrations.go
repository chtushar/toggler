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

// migrationVersion defines the current migration version. This ensures the app
// is always compatible with the version of the database.
const migrationVersion = 1

//go:embed migrations/*.sql
var migrations embed.FS

// buildMigrationClient get the new migrate instance
// source & target connection are needed to close it in the calling function
func buildMigrationClient(db *sql.DB) (*migrate.Migrate, error) {
	// Read the source for the migrations.
	// Our source is the SQL files in the migrations folder
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migration source %w", err)
	}

	// Connect with the target i.e, our postgres DB
	target, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		return nil, fmt.Errorf("failed to read migration target %w", err)
	}

	// Create a new instance of the migration using the defined source and target
	m, err := migrate.NewWithInstance("iofs", source, "postgres", target)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance %w", err)
	}

	return m, nil
}

// upMigrations migrates the postgres schema to the current version
func runUpMigrations(db *sql.DB, logger *log.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Fatal("failed to build client for up migration", err)
		return err
	}

	// Migrate the DB to the current version
	err = m.Migrate(migrationVersion)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		//If the error is present, and the error is not the No change error
		// log the error
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

// runDownMigration migrates the postgres schema from the active migration version to all the way down
func runDownMigration(db *sql.DB, logger *log.Logger) error {
	m, err := buildMigrationClient(db)
	if err != nil {
		logger.Fatal("failed to build client for down migration", err)
		return err
	}

	//applying all down migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		//If the error is present, and the error is not the No change error
		// log the error
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
