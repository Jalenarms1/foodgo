package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Jalenarms1/foodgo/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	dbUrl := os.Getenv("DB_URL")

	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	config.MaxConns = 5

	db.Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Close()

	addr := os.Getenv("API_ADDR")

	server := NewServer(addr)

	log.Fatal(server.Run())
}
