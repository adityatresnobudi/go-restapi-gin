package handler

func (u *userHandler) MapRoutes() {
	u.r.Group("/users").
		POST("/register", u.Create).
		POST("/login", u.Login)
}
