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
	)

	BeforeEach(func() {
		engine = &EngineMock{}
		app = &App{Engine: engine}
		app.SetRoutes()
	})

	Describe("Routes", func() {
		It("/boards", func() {
			Expect(engine).To(Handle("GET").On("/boards/:id").By("github.com/konjoot/treacy/app.Getter"))
			Expect(engine).To(Handle("GET").On("/boards").By("github.com/konjoot/treacy/app.ListGetter"))
			Expect(engine).To(Handle("PUT").On("/boards/:id").By("github.com/konjoot/treacy/app.Updater"))
			Expect(engine).To(Handle("POST").On("/boards").By("github.com/konjoot/treacy/app.Creator"))
			Expect(engine).To(Handle("DELETE").On("/boards/:id").By("github.com/konjoot/treacy/app.Destroyer"))
		})
	})
})
