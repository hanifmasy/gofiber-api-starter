package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang_fiber_api/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables only")
	}

	// Build DSN from env vars
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		// Auto migrate only in dev
		if err := db.AutoMigrate(&models.User{}); err != nil {
			log.Fatal("failed to auto migrate:", err)
		}
		log.Println("✅ Auto migration completed (development mode)")
	} else {
		log.Println("⚠️ Skipping auto migration (production mode)")
	}

	DB = db
	log.Println("Database connected ✅")
}
