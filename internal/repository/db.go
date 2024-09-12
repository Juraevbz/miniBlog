package repository

import (
	"log"
	"mini_blog/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(dsn string) (DB *gorm.DB, err error) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil  {
		log.Fatal("error migrating database", err)
	} 

	return
}