package middleware

import (
	"github.com/Kingpant/golang-clean-architecture-template/internal/interface/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

const ValidatedBodyKey = "validatedBody"

func ValidateBody[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T
		if err := c.BodyParser(&body); err != nil {
			return response.JSONError(c, fiber.StatusBadRequest, "Invalid request body")
		}

		if err := validate.Struct(body); err != nil {
			return response.JSONError(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
		}

		// Set into locals for handler to use
		c.Locals(ValidatedBodyKey, body)
		return c.Next()
	}
}
