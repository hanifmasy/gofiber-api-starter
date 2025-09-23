package response

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	})
}

func Error(c *fiber.Ctx, statusCode int, err interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status: "error",
		Errors: err,
	})
}

func ValidationError(c *fiber.Ctx, err error) error {
	if fe, ok := err.(*fiber.Error); ok {
		return c.Status(fe.Code).JSON(fiber.Map{
			"error": fe.Message,
		})
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": err.Error(),
	})
}
