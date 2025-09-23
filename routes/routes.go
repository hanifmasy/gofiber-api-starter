package routes

import (
	"golang_fiber_api/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes
func SetupRoutes(app *fiber.App, services *services.ServiceRegistry) {
	SetupUserRoutes(app, services.UserService)
	SetupSwaggerRoutes(app)
}
