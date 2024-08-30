package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(dbURL string) error {
	getwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	path := filepath.Join(getwd, "./migrations")

	log.Println(path)

	dbURLFinal := fmt.Sprintf("%s?sslmode=disable",
		dbURL)

	log.Println(dbURLFinal, "migrations...")

	db, err := sql.Open("postgres", dbURLFinal)

	if err != nil {
		return fmt.Errorf("error sql open: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("error driver instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/user-service/migrations/",
		dbURLFinal, driver)

	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
