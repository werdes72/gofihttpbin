package gofihttpbin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func NewApp() *fiber.App {
	app := fiber.New()

	app.Use(favicon.New(favicon.Config{
		File: "./web/static/favicon.ico",
	}))

	httpRoutes(app)
	requestRoutes(app)

	return app
}
