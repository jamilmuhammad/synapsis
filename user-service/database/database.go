package database

import (
	"context"
	"fmt"
	"lib"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDatabase() (*pgxpool.Pool, error) {

	cfg := lib.LoadConfigByFile("./cmd", "config", "yaml")

	// Construct database connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.DB_USER,
		cfg.DB.DB_PASSWORD,
		cfg.DB.DB_HOST,
		cfg.DB.DB_PORT,
		cfg.DB.DB_NAME,
	)

	// Run migrations
	if err := runMigrations(dbURL); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create a connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close()

	return conn, nil
}
