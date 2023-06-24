package gofihttpbin_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGofihttpbin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gofihttpbin Suite")
}
