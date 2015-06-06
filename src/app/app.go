package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunOn(port string) {
	r := &router{gin.Default()}

	r.SetRoutes()

	r.Run(":" + port)
}

type router struct {
	*gin.Engine
}

func (r *router) SetRoutes() {
	fmt.Println("routes are set!")
}
