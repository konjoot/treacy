package app

func (app *App) SetRoutes() {
	app.Engine.GET("/boards", ListGetter)
	app.Engine.GET("/boards/:id", Getter)
	app.Engine.PUT("/boards/:id", Updater)
	app.Engine.POST("/boards", Creator)
	app.Engine.DELETE("/boards/:id", Destroyer)
}
