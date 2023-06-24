package gofihttpbin

import (
	"github.com/gofiber/fiber/v2"
)

func requestRoutes(app fiber.Router) {
	app.Get("/ip", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"origin": c.IP(),
		})
	})

	app.Get("/headers", func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		return c.JSON(fiber.Map{
			"headers": headers,
		})
	})

	app.Get("/user-agent", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user-agent": c.Get("user-agent"),
		})
	})
}
