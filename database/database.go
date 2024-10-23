package database

import (
	"bookstore/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Book{})
}
