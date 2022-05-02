package handlers

import (
	"go-url-shortener/internal/database"
	"go-url-shortener/internal/database/models"
	"go-url-shortener/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

var db = database.InitDatabase()

func MainHandler(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello main!")
	})
}

func AddLink(app *fiber.App) {
	app.Post("/link", func(c *fiber.Ctx) error {
		link := c.Params("link")
		id := helpers.GenerateStringID()
		existing, err := db.GetLink(id)
		if existing != nil || err != nil {
			return c.SendStatus(404)
		}
		DBLink := &models.Link{
			LinkID:   id,
			Original: link,
		}
		return c.JSON(DBLink)
	})
}
