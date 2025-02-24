package handler

func (t *transactionHandler) MapRoutes() {
	t.r.Group("").
		GET("/accounts/:id/transactions", t.GetTransactionById).
		POST("/transfer", t.Create)
}
