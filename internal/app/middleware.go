package app

import "github.com/gofiber/fiber/v2"

func (a *App) AuthMW(c *fiber.Ctx) error {
	if c.Get("X-Allow") == "true" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}