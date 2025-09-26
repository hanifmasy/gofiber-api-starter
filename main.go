package main

// @title           Your API Name
// @version         1.0
// @description     Your API Description
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  your.email@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
import (
	"log"
	"os"

	"golang_fiber_api/pkg/cache"
	"golang_fiber_api/pkg/middleware"

	"github.com/joho/godotenv"

	"golang_fiber_api/database"
	"golang_fiber_api/routes"
	"golang_fiber_api/seeders"
	"golang_fiber_api/services"

	_ "golang_fiber_api/docs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using defaults")
	}

	// Get PORT from .env or fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect DB
	database.ConnectDB()
	// migrate -path migrations -database "postgres://..." up  // manual migration

	db := database.DB
	seeders.Run(db)

	// Initialize service with DB
	serviceRegistry := services.NewServiceRegistry(database.DB)

	cache.Init()

	app := fiber.New()

	app.Use(middleware.CorsMiddleware())

	routes.SetupRoutes(app, serviceRegistry)

	log.Fatal(app.Listen(":" + port))
}
