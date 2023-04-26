package app

func (a *App) InitRoutes() {
	a.router.Get("/", a.Index)
	a.router.Get("/login", a.Login)
	priv := a.router.Group("/priv", a.AuthMW)
	priv.Get("/", a.Priv)
}
