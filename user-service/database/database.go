package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InitDatabase(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {

	log.Println(dbURL, "db...")

	// Create a connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, config)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer conn.Close()

	return conn, nil
}
