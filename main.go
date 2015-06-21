package main

import (
	"github.com/gin-gonic/gin"
	treacy "github.com/konjoot/treacy/app"
)

func main() {
	app := &treacy.App{Engine: gin.Default()}

	app.RunOn("8080")
}
