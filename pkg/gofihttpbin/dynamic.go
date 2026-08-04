package gofihttpbin

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func dynamicRoutes(app fiber.Router) {
	app.Get("/base64/:value", func(c fiber.Ctx) error {
		value := c.Params("value")
		decoded, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Cannot decode given data")
		}
		return c.SendString(string(decoded))
	})

	app.Get("/bytes/:n", func(c fiber.Ctx) error {
		var r *rand.Rand
		seed := int64(fiber.Query[int](c, "seed"))
		n, _ := fiber.Params[int](c, "n", 0), error(nil)

		if seed == 0 {
			r = rand.New(rand.NewSource(time.Now().UnixNano()))
		} else {
			r = rand.New(rand.NewSource(seed))
		}
		blk := make([]byte, n)
		r.Read(blk)

		return c.SendStream(bytes.NewReader(blk))
	})

	app.All("/delay/:delay", func(c fiber.Ctx) error {
		delay, _ := fiber.Params[int](c, "delay", 0), error(nil)
		if delay > 10 {
			delay = 10
		}
		time.Sleep(time.Duration(delay) * time.Second)
		return c.JSON(httpMapper(c))
	})

	app.Get("/uuid", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"uuid": uuid.NewString(),
		})
	})
}
