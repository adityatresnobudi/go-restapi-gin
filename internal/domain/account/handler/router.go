package handler

func (a *accountHandler) MapRoutes() {
	a.r.Group("/accounts", a.auth.Authentication()).
		GET("", a.GetAll).
		GET("/:id", a.GetOne).
		POST("", a.Create).
		PUT("/:id", a.UpdateById).
		DELETE("/:id", a.DeleteById)
}
