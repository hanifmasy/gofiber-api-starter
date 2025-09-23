package routes

import (
	"golang_fiber_api/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes initializes user-related routes
func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	// Swagger annotation for GetUsers endpoint
	// @Summary Get all users
	// @Description Get a list of all users
	// @Tags users
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.User
	// @Router /users [get]
	userGroup.Get("/", controllers.GetUsers)

	userGroup.Get("/:id", controllers.GetUser)

	// Swagger annotation for CreateUser endpoint
	// @Summary Create a new user
	// @Description Create a new user with provided details
	// @Tags users
	// @Accept json
	// @Produce json
	// @Param user body models.User true "User Info"
	// @Success 201 {object} models.User
	// @Router /users [post]
	userGroup.Post("/", controllers.CreateUser)

	userGroup.Put("/:id", controllers.UpdateUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
}
