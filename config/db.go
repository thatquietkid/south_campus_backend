package config

import (
	"fmt"
	"log"

	"github.com/thatquietkid/south_campus_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=south_campus_user password=12345 dbname=south_campus port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	DB = database
	fmt.Println("✅ Connected to Database")

	// Single AutoMigrate call for all models
	err = DB.AutoMigrate(&models.Course{}, &models.CafeteriaItem{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
}
