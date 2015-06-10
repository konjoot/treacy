package app

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
			Expect(engine).To(Handle("GET").On("/boards/:id").By("app.Getter"))
			Expect(engine).To(Handle("GET").On("/boards").By("app.ListGetter"))
			Expect(engine).To(Handle("PUT").On("/boards/:id").By("app.Updater"))
			Expect(engine).To(Handle("POST").On("/boards").By("app.Creator"))
			Expect(engine).To(Handle("DELETE").On("/boards/:id").By("app.Destroyer"))
		})
	})
})
