package validation

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ParseBody parses the request body into `out` struct.
// - If Content-Type is JSON → strict decoding (reject unknown fields).
// - Otherwise → Fiber's BodyParser (form-data, x-www-form-urlencoded).
func ParseBody(c *fiber.Ctx, out interface{}) error {
	ct := c.Get("Content-Type")

	if strings.Contains(ct, "application/json") {
		decoder := json.NewDecoder(bytes.NewReader(c.Body()))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
		}
	} else {
		if err := c.BodyParser(out); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
		}
	}

	return nil
}
