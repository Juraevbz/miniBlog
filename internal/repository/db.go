package repository

import (
	"log"
	"mini_blog/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Like{},
		&models.Repost{},
		&models.Favorite{},
	)

	if err != nil {
		log.Fatal("error migrating database", err)
	}

	return
}
