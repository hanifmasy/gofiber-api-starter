package city

import (
	"golang_fiber_api/pkg/response"
	services "golang_fiber_api/services/city"

	"github.com/gofiber/fiber/v2"
)

func GetCities(cityService *services.CityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cities, err := cityService.GetCities()
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch cities")
		}
		return response.Success(c, cities, "Cities fetched successfully")
	}
}
