package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Creator(c *gin.Context) {
	if r, ok := c.Get("resource"); ok {
		ri := r.(ResourceIface)

		c.Bind(ri.Form())
		ri.Save()

		c.Header("Location", ri.Url())
		c.Data(http.StatusCreated, gin.MIMEJSON, nil)
	}
}

func Getter(c *gin.Context)     {}
func Updater(c *gin.Context)    {}
func ListGetter(c *gin.Context) {}
func Destroyer(c *gin.Context)  {}
