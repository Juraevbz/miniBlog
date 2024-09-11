package main

import (
	"log"
	"mini_blog/internal/repository"
	"os"
)

func main() {
	dsn := os.Getenv("DSN")
	_, err := repository.DBConnection(dsn)
	if err != nil {
		log.Fatal("failed to initialize DBConnection", err)
	}
}