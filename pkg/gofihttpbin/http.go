package gofihttpbin

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func httpRoutes(app fiber.Router) {
	app.Delete("/delete", func(c *fiber.Ctx) error {
		return c.JSON(httpMapper(c))
	})

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"args":    c.Queries(),
			"headers": c.GetReqHeaders(),
			"origin":  c.IP(),
			"url":     c.BaseURL(),
		})
	})

	app.Patch("/patch", func(c *fiber.Ctx) error {
		return c.JSON(httpMapper(c))
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		return c.JSON(httpMapper(c))
	})

	app.Put("/put", func(c *fiber.Ctx) error {
		return c.JSON(httpMapper(c))
	})
}

func httpMapper(c *fiber.Ctx) map[string]interface{} {
	body := c.Request().Body()
	buffer := make([]byte, len(body))
	copy(buffer, body)

	form, err := c.MultipartForm()
	if err != nil {
		form = &multipart.Form{}
	}

	return fiber.Map{
		"args":    c.Queries(),
		"data":    string(buffer),
		"files":   getAllFiles(form),
		"form":    form,
		"headers": c.GetReqHeaders(),
		"json":    getRequestJson(c),
		"origin":  c.IP(),
		"url":     c.BaseURL() + c.Path(),
	}
}

func getRequestJson(c *fiber.Ctx) interface{} {
	if c.GetRespHeader("Content-Type") == "application/json" {
		return c.JSON(c.Body())
	}
	return nil
}

func getAllFiles(f *multipart.Form) map[string][]*multipart.FileHeader {
	return f.File
}
