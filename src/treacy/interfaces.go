package treacy

import "github.com/gin-gonic/gin"

type Engine interface {
	Run(string) error
	GET(relativePath string, handlers ...gin.HandlerFunc)
	PUT(relativePath string, handlers ...gin.HandlerFunc)
	POST(relativePath string, handlers ...gin.HandlerFunc)
	DELETE(relativePath string, handlers ...gin.HandlerFunc)
}
