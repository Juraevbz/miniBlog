package main

import (
	"log"
	"mini_blog/config"
	"mini_blog/internal/handler"
	"mini_blog/internal/repository"
	"mini_blog/internal/service"
)

func main() {
	// TODO: Logger
	// TODO: Validation in service layer
	// TODO: Swagger docs 
	// TODO: Test
	// TODO: Найти все упоминания id и привести в единый тип INT или UINT
	dsn, err := config.InitConfig()
	if err != nil {
		log.Fatal("error initialize configuration", err)
	}

	db, err := repository.DBConnection(dsn)
	if err != nil {
		log.Fatal("failed to initialize DBConnection")
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	if err = handler.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal("error running server", err)
	}
}
