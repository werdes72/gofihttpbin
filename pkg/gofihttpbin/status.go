package gofihttpbin

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func statusRoutes(app fiber.Router) {
	app.All("/status/:code", func(c *fiber.Ctx) (err error) {
		code := c.Params("code")
		var status int

		if strings.Contains(code, ",") {
			codes := strings.Split(code, ",")
			randomIndex := rand.Intn(len(codes))
			status, err = strconv.Atoi(codes[randomIndex])
			if err != nil {
				return c.Status(400).SendString("Invalid status code")
			}
		} else {
			status, err = strconv.Atoi(code)
			if err != nil {
				return c.Status(400).SendString("Invalid status code")
			}
		}

		return c.SendStatus(status)
	})
}
