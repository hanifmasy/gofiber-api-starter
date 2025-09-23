package routes

import (
	_ "golang_fiber_api/docs"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
	_ "github.com/swaggo/files"
)

// SetupSwaggerRoutes initializes Swagger routes
func SetupSwaggerRoutes(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault) // Serve Swagger UI

	// Swagger route to serve the Swagger JSON spec at /docs/doc.json
	app.Get("/docs/doc.json", swagger.New(swagger.Config{
		URL: "/docs/doc.json", // URL to the Swagger JSON endpoint
	}))
}
