package gofihttpbin

import (
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func NewApp(staticDir string) *fiber.App {
	app := fiber.New()

	app.Use(favicon.New(favicon.Config{
		File: path.Join(staticDir, "favicon.ico"),
	}))

	httpRoutes(app)
	requestRoutes(app)

	return app
}
