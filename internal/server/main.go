package server

import "github.com/gofiber/fiber/v2"

func InitServer() {
	app := fiber.New()

	app.Listen(":8080")
}
