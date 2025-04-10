package gofihttpbin

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewApp(staticDir string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "gofihttpbin",
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
		ErrorHandler:          customErrorHandler,
	})

	if val, exists := os.LookupEnv("GOFI_LOGS"); exists && val == "true" {
		app.Use(logger.New())
	}

	dynamicRoutes(app)
	httpRoutes(app)
	requestRoutes(app)
	statusRoutes(app)

	return app
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	if code == fiber.StatusNotFound {
		return c.Status(code).SendString("404 Not Found")
	}
	return c.Status(code).SendString(err.Error())
}
