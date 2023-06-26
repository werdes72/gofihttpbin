package gofihttpbin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func dynamicRoutes(app fiber.Router) {
	app.Get("/uuid", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"uuid": uuid.NewString(),
		})
	})
}
