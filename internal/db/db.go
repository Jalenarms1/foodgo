package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func SetPool() error {
	dbUrl := os.Getenv("DB_URL")

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return err
	}

	config.MaxConns = 5

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return nil
}
