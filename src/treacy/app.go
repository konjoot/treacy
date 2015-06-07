package treacy

import (
	"fmt"
)

type App struct {
	Engine Engine
}

type Engine interface {
	Run(string) error
}

func (app *App) RunOn(port string) {
	app.SetRoutes()

	app.Engine.Run(":" + port)
}

func (app *App) SetRoutes() {
	fmt.Println("routes are set!")
}
