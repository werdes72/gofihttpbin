package gofihttpbin_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gofiber/fiber/v2"
	"github.com/werdes72/gofihttpbin/pkg/gofihttpbin"
)

var _ = Describe("Request routes", func() {
	app := gofihttpbin.NewApp("../../web/static/")

	It("/ip returns client's IP", func() {
		req := httptest.NewRequest("GET", "/ip", nil)

		res, _ := app.Test(req)
		body, err := io.ReadAll(res.Body)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(body).To(ContainSubstring("origin"))
		Expect(body).To(ContainSubstring("0.0.0.0"))
	})

	It("/headers returns request headers", func() {
		req := httptest.NewRequest("GET", "/headers", nil)
		req.Header.Set("Accept-Language", "en-us")
		req.Header.Set("User-Agent", "test")

		res, _ := app.Test(req)
		var resJSON map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON).To(HaveKey("headers"))
		Expect(resJSON["headers"]).To(HaveKeyWithValue("Accept-Language", "en-us"))
		Expect(resJSON["headers"]).To(HaveKeyWithValue("User-Agent", "test"))
	})

	It("/user-agent returns request user agent", func() {
		req := httptest.NewRequest("GET", "/user-agent", nil)
		req.Header.Set("User-Agent", "test")

		res, _ := app.Test(req)
		var resJSON map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON).To(HaveKey("user-agent"))
		Expect(resJSON["user-agent"]).To(Equal("test"))
	})
})
