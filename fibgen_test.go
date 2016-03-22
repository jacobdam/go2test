package go2test_test

import (
	. "github.com/jacobdam/go2test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewFibGen", func() {
	It("should return fibonacci generator", func() {
		f := NewFibGen()
		Expect(f()).To(Equal(0))
		Expect(f()).To(Equal(1))
		Expect(f()).To(Equal(1))
		Expect(f()).To(Equal(2))
		Expect(f()).To(Equal(3))
		Expect(f()).To(Equal(5))
	})
})
