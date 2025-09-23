package controllers

import (
	"golang_fiber_api/database"
	"golang_fiber_api/dtos"
	"golang_fiber_api/models"
	"golang_fiber_api/pkg/response"
	"golang_fiber_api/pkg/validation"
	"golang_fiber_api/services"
	"strconv"

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
func GetUsers(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))

		users, err := userService.GetUsers(page, limit)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch users")
		}

		return response.Success(c, users, "Users fetched successfully")
	}
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
func GetUser(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		var user models.User
		result := database.DB.First(&user, id)
		if result.Error != nil {
			return response.Error(c, fiber.ErrNotFound.Code, "User not found")
		}

		var data = dtos.UserResponseDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		return response.Success(c, data, "User fetched successfully")
	}
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
func CreateUser(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto dtos.CreateUserDTO
		if err := validation.ParseBody(c, &dto); err != nil {
			return response.ValidationError(c, err)
		}
		if errs := validation.ValidateStruct(dto); errs != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		// hash password
		hashedPassword, err := HashPassword(dto.Password)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to hash password")
		}
		dto.Password = hashedPassword

		createdUser, err := userService.CreateUser(dto)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create user")
		}

		return response.Success(c, createdUser, "User created successfully")
	}
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
func UpdateUser(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		// fetch existing user
		user, err := userService.GetUserByID(id)
		if err != nil {
			return response.Error(c, fiber.ErrNotFound.Code, "User not found")
		}

		var dto dtos.UpdateUserDTO
		if err := validation.ParseBody(c, &dto); err != nil {
			return response.ValidationError(c, err)
		}
		if errs := validation.ValidateStruct(dto); errs != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		// apply updates
		if dto.Name != nil {
			user.Name = *dto.Name
		}
		if dto.Email != nil {
			user.Email = *dto.Email
		}
		if dto.Password != nil {
			hashedPassword, err := HashPassword(*dto.Password)
			if err != nil {
				return response.Error(c, fiber.StatusInternalServerError, "Failed to hash password")
			}
			user.Password = hashedPassword
		}

		// update in DB
		if err := userService.UpdateUser(user); err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update user")
		}

		return response.Success(c, dtos.ToUserResponseDTO(*user), "User updated successfully")
	}
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
func DeleteUser(userService *services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
		}

		// call service directly with id
		deletedUser, err := userService.DeleteUser(id)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to delete user")
		}

		return response.Success(c, deletedUser, "User deleted successfully")
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
