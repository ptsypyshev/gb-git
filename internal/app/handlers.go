package app

import "github.com/gofiber/fiber/v2"

func (a *App) Index(c *fiber.Ctx) error {
	return c.SendString("Hello world!")
}

func (a *App) Priv(c *fiber.Ctx) error {
	return c.SendString("Private content")
}