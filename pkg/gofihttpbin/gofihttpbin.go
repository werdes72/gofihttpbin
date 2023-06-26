package gofihttpbin

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewApp(staticDir string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "gofihttpbin",
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
		ErrorHandler:          customErrorHandler,
	})

	dynamicRoutes(app)
	httpRoutes(app)
	requestRoutes(app)

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
