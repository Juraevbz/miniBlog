package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database", err)
	}

	// err = db.AutoMigrate()
	// if err != nil  {
	// 	log.Fatal("error migrating database", err)
	// } 

	return
}