package box_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Box Suite")
}
