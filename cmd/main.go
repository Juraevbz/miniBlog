package main

import (
	"mini_blog/config"
	"mini_blog/internal/handler"
	"mini_blog/internal/repository"
	"mini_blog/internal/service"
	"mini_blog/pkg/glog"
)

func main() {
	logger := glog.NewLogger()

	dsn, err := config.InitConfig(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("error initializing configuration")
	}

	db, err := repository.DBConnection(dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize DBConnection")
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository, logger)
	handler := handler.NewHandler(service)

	logger.Info().Msg("server listening on port 8080")
	if err = handler.Run("8080", handler.InitRoutes()); err != nil {
		logger.Fatal().Err(err).Msg("error running server")
	}
}
