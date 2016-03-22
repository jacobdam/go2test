package go2test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGo2test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go2test Suite")
}
