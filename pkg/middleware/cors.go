package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CorsMiddleware returns Fiber CORS middleware configured from .env
func CorsMiddleware() fiber.Handler {
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	allowedHeaders := os.Getenv("ALLOWED_HEADERS")
	if allowedHeaders == "" {
		allowedHeaders = "Content-Type,Authorization"
	}

	allowedMethods := os.Getenv("ALLOWED_METHODS")
	if allowedMethods == "" {
		allowedMethods = "GET,POST,DELETE,PUT,OPTIONS"
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     strings.ReplaceAll(allowedHeaders, " ", ""),
		AllowMethods:     strings.ReplaceAll(allowedMethods, " ", ""),
		AllowCredentials: os.Getenv("ALLOW_CREDENTIALS") == "true",
	})

}
