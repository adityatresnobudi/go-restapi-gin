package handler

func (t *transactionHandler) MapRoutes() {
	t.r.Group("", t.auth.Authentication()).
		GET("/accounts/:id/transactions", t.GetTransactionById).
		POST("/transfer", t.Create)
}
