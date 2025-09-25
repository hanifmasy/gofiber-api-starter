package user

import (
	"golang_fiber_api/dtos"
	"golang_fiber_api/pkg/response"
	"golang_fiber_api/pkg/validation"
	services "golang_fiber_api/services/user"

	"github.com/gofiber/fiber/v2"
)

func Signup(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto dtos.UserSignupDTO
		if err := validation.ParseBody(c, &dto); err != nil {
			return response.ValidationError(c, err)
		}
		if errs := validation.ValidateStruct(dto); errs != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		user, err := authService.Signup(dto)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		}
		return response.Success(c, user, "User registered successfully")
	}
}

func Signin(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dto dtos.UserSigninDTO
		if err := validation.ParseBody(c, &dto); err != nil {
			return response.ValidationError(c, err)
		}
		if errs := validation.ValidateStruct(dto); errs != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
		}

		token, err := authService.Signin(dto)
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, err.Error())
		}

		return response.Success(c, fiber.Map{"token": token}, "Login successful")
	}
}

func Signout(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authService.Signout()
		return response.Success(c, nil, "Logout successful")
	}
}
