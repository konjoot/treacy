package treacy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var (
		app    *App
		engine *EngineMock
	)

	BeforeEach(func() {
		engine = &EngineMock{}
		app = &App{Engine: engine}
		app.SetRoutes()
	})

	Describe("Routes", func() {
		It("POST /boards", func() {
			Expect(engine).To(Handle("POST").On("/boards").By(BoardsCreator))
		})
	})
})
