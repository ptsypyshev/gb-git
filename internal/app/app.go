package app

import "github.com/gofiber/fiber/v2"

type App struct {
	router *fiber.App
}

func NewApp() *App {
	r := fiber.New()
	return &App{
		router: r,
	}
}

func (a *App) Run()  {
	a.router.Listen(":8080")
}