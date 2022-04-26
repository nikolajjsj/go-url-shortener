package handlers

import "github.com/gofiber/fiber/v2"

func MainHandler(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello main!")
	})
}
