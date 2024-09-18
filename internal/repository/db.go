package repository

import (
	"log"
	"mini_blog/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}

func DBConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database", err)
	}

	err = db.AutoMigrate(
		//&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Like{},
		&models.Repost{},
	)

	if err != nil {
		log.Fatal("error migrating database", err)
	}

	return
}
