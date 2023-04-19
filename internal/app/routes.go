package app

func (a *App) InitRoutes()  {
	a.router.Get("/", a.Index)
}