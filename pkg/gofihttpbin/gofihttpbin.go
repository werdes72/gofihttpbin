package gofihttpbin

import (
	"github.com/gofiber/fiber/v2"
)

func NewApp(staticDir string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "gofihttpbin",
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
	})

	dynamicRoutes(app)
	httpRoutes(app)
	requestRoutes(app)

	return app
}
