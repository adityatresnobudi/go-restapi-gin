package handler

func (a *accountHandler) MapRoutes() {
	a.r.Group("/accounts", a.auth.Authentication(), a.auth.AdminAuthorization()).
		GET("", a.GetAll).
		POST("", a.Create).
		PUT("/:id", a.UpdateById).
		DELETE("/:id", a.DeleteById)
	a.r.Group("/accounts").
		GET("/:id", a.GetOne)
}
