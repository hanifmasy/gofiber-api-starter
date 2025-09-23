package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes
func SetupRoutes(app *fiber.App) {
	SetupUserRoutes(app)
	SetupSwaggerRoutes(app)
}
