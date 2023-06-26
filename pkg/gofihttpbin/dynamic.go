package gofihttpbin

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func dynamicRoutes(app fiber.Router) {
	app.Get("/base64/:value", func(c *fiber.Ctx) error {
		value := c.Params("value")
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Cannot decode given data")
		}
		return c.SendString(string(decoded))
	})

	app.Get("/uuid", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"uuid": uuid.NewString(),
		})
	})
}
