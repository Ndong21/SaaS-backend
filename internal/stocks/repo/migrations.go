package repo

import (
	"errors"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate function applies migrations to the database
func MigrateUp(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file://"+absPath, dbURL,
	)

	if err != nil {
		return err
	}

	defer func() {
		err, _ := m.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//Apply migrations
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// MigrateDown  function rolls back migrations from the database
func MigrateDown(dbURL string, migrationsPath string) error {
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return err
	}

	// Create a new migration instance with the absolute path
	m, err := migrate.New(
		"file:/"+absPath,
		dbURL,
	)
	if err != nil {
		return err
	}

	defer func() {
		err, _ = m.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Apply migrations
	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
