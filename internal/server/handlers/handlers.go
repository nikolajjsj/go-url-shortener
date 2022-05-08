package handlers

import (
	"go-url-shortener/internal/database"
	"go-url-shortener/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App, db *database.DB) {
	mainHandler(app, db)
	addLink(app, db)
	getLink(app, db)
}

func mainHandler(app *fiber.App, db *database.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello main!")
	})
}

func addLink(app *fiber.App, db *database.DB) {
	app.Post("/link/:link", func(c *fiber.Ctx) error {
		link := c.Params("link")
		id := helpers.GenerateStringID()
		existing, _ := db.GetLink(id)
		if existing != nil {
			return c.SendStatus(404)
		}
		dbLink, err := db.CreateLink(id, link)
		if err != nil {
			return c.SendStatus(404)
		}
		return c.JSON(dbLink)
	})
}

func getLink(app *fiber.App, db *database.DB) {
	app.Get("/*", func(c *fiber.Ctx) error {
		link := c.OriginalURL()
		linkID := link[1:]
		existing, err := db.GetLink(linkID)
		if err != nil || existing == nil {
			return c.SendStatus(404)
		}
		return c.Redirect("http://" + existing.Original)
	})
}
