package response

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	OK   bool `json:"ok"`
	Data any  `json:"data,omitempty"`
}

type ErrorResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func JSONOK(c *fiber.Ctx, data any) error {
	return c.JSON(BaseResponse{OK: true, Data: data})
}

func JSONError(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(ErrorResponse{OK: false, Error: msg})
}
