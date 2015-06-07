package main

import (
	"github.com/gin-gonic/gin"
	"treacy"
)

func main() {
	app := &treacy.App{Engine: gin.Default()}

	app.RunOn("8080")
}
