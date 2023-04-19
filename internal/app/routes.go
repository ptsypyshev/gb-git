package app

func (a *App) InitRoutes()  {
	a.router.Get("/", a.Index)
	priv := a.router.Group("/priv", a.AuthMW)
	priv.Get("/", a.Priv)
}