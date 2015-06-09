package treacy

type App struct {
	Engine Engine
}

func (app *App) RunOn(port string) {
	app.SetRoutes()

	app.Engine.Run(":" + port)
}
