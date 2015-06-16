package app

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Handlers", func() {
	var (
		router   *gin.Engine
		response *httptest.ResponseRecorder
		request  *http.Request
		resource *ResourceMock
		body     map[string]string
		cType    string
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		response = httptest.NewRecorder()
		cType = gin.MIMEJSON
	})

	Describe("Creator", func() {
		JustBeforeEach(func() {
			router.POST("/tests", Creator)
			jsBody, _ := json.Marshal(body)
			request, _ = http.NewRequest("POST", "/tests", bytes.NewBuffer(jsBody))
			request.Header.Add("Content-Type", cType)

			router.ServeHTTP(response, request)
		})

		Context("on success", func() {
			BeforeEach(func() {
				resource = Resource()
				router.Use(func(c *gin.Context) {
					c.Set("resource", resource)
				})
				body = map[string]string{
					"Name": "test",
					"Desc": "testDesc"}
			})

			It("should bind new resource", func() {
				Expect(resource).To(BindedWith(body))
			})
			It("should create new resource", func() {
				Expect(resource).To(BeCreated())
			})
			It("should return status 201", func() {
				Expect(response.Code).To(Equal(201))
			})
			It("should return Location header with a resource link", func() {
				Expect(response).To(HaveHeader("Location").With(resource.Url()))
			})
			It("should return empty body", func() {
				Expect(response.Body.Len()).To(BeZero())
			})
		})

		Context("when conflict", func() {
			It("should not create new resource", func() {})
			It("should return status 409", func() {})
			It("should return body with error desc", func() {})
		})

		Context("when invalid params", func() {
			It("should not create new resource", func() {})
			It("should return status 400", func() {})
			It("should return body with error desc", func() {})
		})
	})
})
