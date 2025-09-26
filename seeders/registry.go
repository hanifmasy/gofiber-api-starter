package seeders

import (
	"log"
	"os"

	"gorm.io/gorm"
)

// SeederFunc defines the function signature for a seeder
type SeederFunc func(db *gorm.DB)

// Registry holds all seeders to be executed
var Registry = []SeederFunc{
	SeedCities,
	// Add more seeders here
}

// Run executes all registered seeders (only in development)
func Run(db *gorm.DB) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "development" {
		log.Println("⚠️ Skipping seeders (not development mode)")
		return
	}

	log.Println("🚀 Running seeders...")
	for _, seeder := range Registry {
		seeder(db)
	}
	log.Println("✅ Seeders completed")
}
