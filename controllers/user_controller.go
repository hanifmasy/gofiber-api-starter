package controllers

import (
	"golang_fiber_api/database"
	"golang_fiber_api/dtos"
	"golang_fiber_api/models"
	"golang_fiber_api/pkg/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers godoc
// @Summary      Get all users
// @Description  Get a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      400  {object}  error
// @Router       /users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Get a single user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  error
// @Router       /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "User Data"
// @Success      201   {object}  models.User
// @Failure      400   {object}  error
// @Router       /users [post]
func CreateUser(c *fiber.Ctx) error {
	var dto dtos.CreateUserDTO

	if err := validation.ParseBody(c, &dto); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Hash password
	hashedPassword, err := HashPassword(dto.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	database.DB.Create(&user)
	return c.Status(201).JSON(user)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update an existing user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      models.User true  "User Data"
// @Success      200   {object}  models.User
// @Failure      404   {object}  error
// @Router       /users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	var dto dtos.UpdateUserDTO

	if err := validation.ParseBody(c, &dto); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if dto.Name != nil {
		user.Name = *dto.Name
	}
	if dto.Email != nil {
		user.Email = *dto.Email
	}
	if dto.Password != nil {
		hashedPassword, err := HashPassword(*dto.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = hashedPassword
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete an existing user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {string}  string "User deleted"
// @Failure      404  {object}  error
// @Router       /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	database.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted"})
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
