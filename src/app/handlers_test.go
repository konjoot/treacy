package app

import (
	"bytes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Handlers", func() {
	var (
		router   *gin.Engine
		response *httptest.ResponseRecorder
		request  *http.Request
		body     *bytes.Buffer
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		response = httptest.NewRecorder()
	})

	Describe("Creator", func() {
		BeforeEach(func() {
			router.POST("/boards", Creator)
		})

		Context("success", func() {
			BeforeEach(func() {
				body = bytes.NewBufferString("")
				request, _ = http.NewRequest("POST", "/boards", body)
				request.Header.Add("Content-Type", gin.MIMEJSON)
				router.ServeHTTP(response, request)
			})

			It("should create new resource", func() {})
			It("should return status 201", func() {
				Expect(response.Code).To(Equal(201))
			})
			It("should return Location header with a resource link", func() {})
		})

		Context("conflict", func() {
			It("should not create new resource", func() {})
			It("should return status 409", func() {})
			It("should return body with error desc", func() {})
		})

		Context("invalid params", func() {
			It("should not create new resource", func() {})
			It("should return status 400", func() {})
			It("should return body with error desc", func() {})
		})
	})
})
