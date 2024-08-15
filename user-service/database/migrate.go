package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(dbURL string) error {
	getwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	path := "file://" + filepath.Join(getwd, "user_service/migrations")

	m, err := migrate.New(
		path,
		dbURL)

	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
