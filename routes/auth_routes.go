package routes

import (
	controllers "golang_fiber_api/controllers/user"
	services "golang_fiber_api/services/user"

	"github.com/gofiber/fiber/v2"
)

func SetupUserAuthRoutes(app *fiber.App, authService *services.AuthService) {
	auth := app.Group("/auth")

	auth.Post("/signup", controllers.Signup(authService))
	auth.Post("/signin", controllers.Signin(authService))
	auth.Post("/signout", controllers.Signout(authService))
}
