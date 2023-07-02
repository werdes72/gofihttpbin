package gofihttpbin

import (
	"encoding/json"
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
			"url":     c.BaseURL() + c.Path(),
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
	//	body := c.Request().Body()
	//	buffer := make([]byte, len(body))
	//	copy(buffer, body)
	body := c.Request().Body()

	form, err := c.MultipartForm()
	if err != nil {
		form = &multipart.Form{}
	}

	return fiber.Map{
		"args":    c.Queries(),
		"data":    string(body),
		"files":   getAllFiles(form),
		"form":    form,
		"headers": c.GetReqHeaders(),
		"json":    getRequestJson(c, body),
		"origin":  c.IP(),
		"url":     c.BaseURL() + c.Path(),
	}
}

func getRequestJson(c *fiber.Ctx, b []byte) interface{} {
	header, exists := c.GetReqHeaders()["Content-Type"]
	if exists && header == "application/json" {
		j := map[string]interface{}{}
		err := json.Unmarshal(b, &j)
		if err == nil {
			return j
		}
	}
	return nil
}

func getAllFiles(f *multipart.Form) map[string][]*multipart.FileHeader {
	return f.File
}
