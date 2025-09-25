package routes

import (
	services "golang_fiber_api/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes
func SetupRoutes(app *fiber.App, services *services.ServiceRegistry) {
	SetupUserAuthRoutes(app, services.UserAuthService)
	SetupUserRoutes(app, services.UserService)
	SetupSwaggerRoutes(app)
}
