package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/domain/transaction/service"
	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/internal/middleware/auth"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	r       *gin.Engine
	ctx     context.Context
	auth    auth.AuthMiddleware
	service service.TransactionService
}

func NewTransactionHandler(
	r *gin.Engine,
	ctx context.Context,
	auth auth.AuthMiddleware,
	service service.TransactionService,
) *transactionHandler {
	return &transactionHandler{
		r:       r,
		ctx:     ctx,
		auth:    auth,
		service: service,
	}
}

// @Summary Get All Transaction By ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object}  GetTransactionByIdResponseDTO
// @Router /accounts/{id}/transactions [get]
func (t *transactionHandler) GetTransactionById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	result, err := t.service.GetTransactionById(ctx, id)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create Transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param requestBody body CreateTransactionRequestDTO true "Request Body"
// @Success 201 {object} CreateTransactionResponseDTO
// @Router /transactions [post]
func (t *transactionHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.CreateTransactionRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, errData := t.service.Create(ctx, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusCreated, result)
}
