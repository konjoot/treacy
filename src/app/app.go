package app

type App struct {
	Engine EngineIface
}

func (app *App) RunOn(port string) {
	app.SetRoutes()

	app.Engine.Run(":" + port)
}
