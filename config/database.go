package database

import (
	"go-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=farhan password=farhan0212 dbname=golang_backend port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	log.Println("Database connected successfully")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database", err)
	}
	log.Println("Database migration completed")
}
