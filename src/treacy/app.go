package treacy

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine Engine
}

type Engine interface {
	Run(string) error
	POST(relativePath string, handlers ...gin.HandlerFunc)
}

func (app *App) RunOn(port string) {
	app.SetRoutes()

	app.Engine.Run(":" + port)
}

func (app *App) SetRoutes() {
	app.Engine.POST("/boards", Creator)
}

func Creator(c *gin.Context) {}
