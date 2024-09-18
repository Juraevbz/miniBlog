package main

import (
	"log"
	"mini_blog/config"
	"mini_blog/internal/handler"
	"mini_blog/internal/repository"
	"mini_blog/internal/service"
)

func main() {

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

	log.Println("Server is listening...")
	if err = handler.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal("error running server", err)
	}

}
