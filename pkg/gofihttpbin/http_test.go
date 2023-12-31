package gofihttpbin_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gofiber/fiber/v2"
	"github.com/werdes72/gofihttpbin/pkg/gofihttpbin"
)

var _ = Describe("HTTP routes", func() {
	app := gofihttpbin.NewApp("../../web/static/")

	tests := map[string]string{
		"/delete": "DELETE",
		"/get":    "GET",
		"/patch":  "PATCH",
		"/post":   "POST",
		"/put":    "PUT",
	}

	for path, method := range tests {
		path := path
		method := method
		It(fmt.Sprintf("%s returns 200", path), func() {
			req := httptest.NewRequest(method, fmt.Sprintf("%s?a=b&c=d", path), strings.NewReader("Body"))

			res, _ := app.Test(req)
			var resJSON map[string]interface{}
			err := json.NewDecoder(res.Body).Decode(&resJSON)

			Expect(err).NotTo(HaveOccurred())
			Expect(res.StatusCode).To(Equal(fiber.StatusOK))
			Expect(resJSON["args"]).To(HaveKeyWithValue("a", "b"))
			Expect(resJSON["args"]).To(HaveKeyWithValue("c", "d"))
			Expect(resJSON["url"]).To(ContainSubstring("http://example.com" + path))
			Expect(resJSON["origin"]).To(ContainSubstring("0.0.0.0"))
			if method != "GET" {
				Expect(resJSON["data"]).To(Equal("Body"))
			}
		})
	}

	It("/delete return JSON body", func() {
		req := httptest.NewRequest("DELETE", "/delete", strings.NewReader(`{"Key": "Val"}`))
		req.Header.Add("Content-Type", "application/json")

		res, _ := app.Test(req)
		var resJSON map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&resJSON)

		Expect(err).NotTo(HaveOccurred())
		Expect(res.StatusCode).To(Equal(fiber.StatusOK))
		Expect(resJSON["json"]).To(HaveKeyWithValue("Key", "Val"))
		Expect(resJSON["url"]).To(ContainSubstring("http://example.com/delete"))
		Expect(resJSON["origin"]).To(ContainSubstring("0.0.0.0"))
	})
})
