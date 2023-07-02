package gofihttpbin_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/werdes72/gofihttpbin/pkg/gofihttpbin"
)

var _ = Describe("Dynamic routes", func() {
	app := gofihttpbin.NewApp("../../web/static/")

	It("/base64/ returns 404", func() {
		req := httptest.NewRequest("GET", "/base64/", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusNotFound))
		Expect(body).To(ContainSubstring("404 Not Found"))
	})

	It("/base64/SFRUUEJJTiBpcyBhd2Vzb21l returns HTTPBIN is awesome", func() {
		req := httptest.NewRequest("GET", "/base64/SFRUUEJJTiBpcyBhd2Vzb21l", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(body).To(ContainSubstring("HTTPBIN is awesome"))
	})

	It("/base64/asd returns an error", func() {
		req := httptest.NewRequest("GET", "/base64/asd", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
		Expect(body).To(ContainSubstring("Cannot decode given data"))
	})

	It("/bytes/0 returns 0 random bytes", func() {
		req := httptest.NewRequest("GET", "/bytes/0", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(len(body)).To(Equal(0))
	})

	It("/bytes/16 returns 16 random bytes", func() {
		req := httptest.NewRequest("GET", "/bytes/16", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(len(body)).To(Equal(16))
	})

	It("/bytes/16 with seed returns the same 16 random bytes", func() {
		var data [][]byte

		for i := 0; i < 2; i++ {
			req := httptest.NewRequest("GET", "/bytes/16?seed=123456", nil)

			res, _ := app.Test(req)
			body, err := io.ReadAll(res.Body)
			data = append(data, body)

			Expect(err).NotTo(HaveOccurred())
			Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		}

		Expect(len(data)).To(Equal(2))
		Expect(len(data[0])).To(Equal(16))
		Expect(len(data[1])).To(Equal(16))
		Expect(data[0]).To(Equal(data[1]))
	})

	It("/delay/1 returns 200 after 1s", func() {
		req := httptest.NewRequest("GET", "/delay/1", nil)

		res, _ := app.Test(req, 1500)
		var resJSON map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON["url"]).To(ContainSubstring("http://example.com/delay/1"))
		Expect(resJSON["origin"]).To(ContainSubstring("0.0.0.0"))
	})

	It("/delay/15 returns 200 after 10s", func() {
		req := httptest.NewRequest("GET", "/delay/15", nil)

		res, _ := app.Test(req, 10500)
		var resJSON map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON["url"]).To(ContainSubstring("http://example.com/delay/15"))
		Expect(resJSON["origin"]).To(ContainSubstring("0.0.0.0"))
	})

	It("/uuid returns uuid", func() {
		req := httptest.NewRequest("GET", "/uuid", nil)

		res, _ := app.Test(req)
		var resJSON map[string]string
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON).To(HaveKey("uuid"))
		_, err = uuid.Parse(resJSON["uuid"])
		Expect(err).NotTo(HaveOccurred())
	})
})
