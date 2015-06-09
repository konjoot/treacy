package treacy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "mocks"
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
		It("/boards", func() {
			Expect(engine).To(Handle("GET").On("/boards/:id").By("treacy.Getter"))
			Expect(engine).To(Handle("GET").On("/boards").By("treacy.ListGetter"))
			Expect(engine).To(Handle("PUT").On("/boards/:id").By("treacy.Updater"))
			Expect(engine).To(Handle("POST").On("/boards").By("treacy.Creator"))
			Expect(engine).To(Handle("DELETE").On("/boards/:id").By("treacy.Destroyer"))
		})
	})
})
