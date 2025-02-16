package core

import (
	"github.com/gofiber/fiber/v2"
)

func WithSuccess(c *fiber.Ctx, msg string, data fiber.Map) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": msg,
		"data":    data,
	})
}

func WithError(c *fiber.Ctx, msg string, statusCode int) error {
	return c.Status(statusCode).
		JSON(fiber.Map{
			"status":  false,
			"message": msg,
		})
}
