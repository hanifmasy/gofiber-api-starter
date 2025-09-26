package seeders

import (
	"log"

	"golang_fiber_api/models"

	"gorm.io/gorm"
)

func SeedCities(db *gorm.DB) {
	cities := []models.City{
		{Name: "Jakarta"},
		{Name: "Bandung"},
		{Name: "Surabaya"},
		{Name: "Medan"},
		{Name: "Yogyakarta"},
	}

	for _, city := range cities {
		var existing models.City
		if err := db.Where("name = ?", city.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&city).Error; err != nil {
				log.Printf("❌ Failed to seed city %s: %v\n", city.Name, err)
			} else {
				log.Printf("✅ Seeded city: %s\n", city.Name)
			}
		}
	}
}
