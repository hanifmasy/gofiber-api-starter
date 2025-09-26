package city

import (
	"encoding/json"
	"log"
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
	const cacheKey = "cities:list"

	// 1. Try Redis cache
	cached, err := cache.Rdb.Get(cache.Ctx, cacheKey).Result()
	if err == nil {
		var cities []models.City
		if jsonErr := json.Unmarshal([]byte(cached), &cities); jsonErr == nil {
			log.Println("📦 Cache hit for cities")
			return cities, nil
		}
	}

	// 2. Fetch from DB if not cached
	var cities []models.City
	if err := s.db.Find(&cities).Error; err != nil {
		return nil, err
	}

	// 3. Store in Redis with TTL
	jsonData, _ := json.Marshal(cities)
	cache.Rdb.Set(cache.Ctx, cacheKey, jsonData, 1*time.Hour)

	log.Println("💾 Cities cached in Redis")
	return cities, nil
}
