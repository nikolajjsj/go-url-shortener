package database

import (
	"log"

	"go-url-shortener/internal/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	client *gorm.DB
}

func InitDatabase() *DB {
	client, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db := &DB{
		client: client,
	}

	err = client.AutoMigrate(&models.Link{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
