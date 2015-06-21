package app

import (
	. "github.com/konjoot/treacy/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var (
		app    *App
		engine *EngineMock
		port   string
	)

	BeforeEach(func() {
		port = "8080"
		engine = &EngineMock{}
		app = &App{Engine: engine}
	})

	Describe("RunOn", func() {
		It("should run engine on specified port", func() {
			Expect(engine).NotTo(BeRunning())
			Expect(engine.Port()).To(BeZero())

			app.RunOn(port)

			Expect(engine).To(BeRunning())
			Expect(engine.Port()).To(Equal(":" + port))
		})
	})
})
