package routes

import (
	controllers "golang_fiber_api/controllers/city"
	"golang_fiber_api/pkg/middleware"
	services "golang_fiber_api/services/city"

	"github.com/gofiber/fiber/v2"
)

func SetupCityRoutes(app *fiber.App, cityService *services.CityService) {
	cityGroup := app.Group("/cities")
	protected := cityGroup.Use(middleware.JWTMiddleware())

	protected.Get("/", controllers.GetCities(cityService))
}
