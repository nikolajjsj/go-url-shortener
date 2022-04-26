package server

import (
	"go-url-shortener/internal/server/handlers"
	"go-url-shortener/internal/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	app := fiber.New()

	// Middleware
	middleware.LoggerMiddleware(app)

	// Handlers
	handlers.MainHandler(app)

	app.Listen(":8080")
}
