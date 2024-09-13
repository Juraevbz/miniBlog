package main

import (
	"log"
	"mini_blog/internal/handler"
	"mini_blog/internal/repository"
	"mini_blog/internal/service"
	"os"
)

func main() {
	dsn := os.Getenv("DSN")
	db, err := repository.DBConnection(dsn)
	if err != nil {
		log.Fatal("failed to initialize DBConnection", err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	if err = handler.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal("error running server", err)
	}

}
