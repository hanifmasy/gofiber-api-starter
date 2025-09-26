package city

import (
	"time"

	"golang_fiber_api/models"
	"golang_fiber_api/pkg/cache"

	"gorm.io/gorm"
)

type CityService struct {
	db *gorm.DB
}

func NewCityService(db *gorm.DB) *CityService {
	return &CityService{db: db}
}

func (s *CityService) GetCities() ([]models.City, error) {
	// Fetch from DB if not cached
	var cities []models.City

	err := cache.GetOrSetCache(
		"cities:all",
		1*time.Hour,
		&cities,
		func() (*[]models.City, error) {
			var dbCities []models.City
			if err := s.db.Find(&dbCities).Error; err != nil {
				return nil, err
			}
			return &dbCities, nil
		},
	)

	return cities, err
}
