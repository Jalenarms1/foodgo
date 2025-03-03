package main

import (
	"log"
	"os"

	"github.com/Jalenarms1/foodgo/internal/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := db.SetPool()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Pool.Close()

	addr := os.Getenv("API_ADDR")

	server := NewServer(addr)

	log.Fatal(server.Run())
}
