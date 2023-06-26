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

	It("/base64/asd returns ", func() {
		req := httptest.NewRequest("GET", "/base64/asd", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusBadRequest))
		Expect(body).To(ContainSubstring("Cannot decode given data"))
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
