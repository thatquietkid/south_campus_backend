package config

import (
	"fmt"
	"log"
	"os"

	"github.com/thatquietkid/south_campus_backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	DB = database
	fmt.Println("✅ Connected to Database")

	if os.Getenv("APP_ENV") != "production" {
		err = DB.AutoMigrate(&models.Course{}, &models.CafeteriaItem{}, &models.User{})
		if err != nil {
			log.Fatalf("❌ AutoMigrate failed: %v", err)
		}
	}

	return DB
}
