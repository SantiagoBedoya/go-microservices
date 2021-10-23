package middlewares

import "github.com/gofiber/fiber/v2"

func IsAuthenticated(c *fiber.Ctx) error {
	return c.JSON("auth middleware")
}
