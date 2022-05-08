package server

import (
	"go-url-shortener/internal/database"
	"go-url-shortener/internal/server/handlers"
	"go-url-shortener/internal/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitServer() {
	// Init database
	db := database.InitDatabase()

	// Start fiber router
	app := fiber.New()

	// Middleware
	middleware.LoggerMiddleware(app)

	// Handlers
	handlers.InitHandlers(app, db)

	app.Listen(":8080")
}
