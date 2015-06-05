package app

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	It("should work", func() {
		Expect("work").To(Equal("work"))
	})
})
