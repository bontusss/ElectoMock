package database

import (
	"electomock/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	*gorm.DB
}

func NewPostgresDB(url string) *Database {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//	Auto migrate schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}

	return &Database{db}
}

func (db *Database) Close() {
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Printf("Error closing database: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}
