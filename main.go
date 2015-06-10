package main

import (
	treacy "app"
	"github.com/gin-gonic/gin"
)

func main() {
	app := &treacy.App{Engine: gin.Default()}

	app.RunOn("8080")
}
