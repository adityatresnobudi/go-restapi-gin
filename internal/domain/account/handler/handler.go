package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/domain/account/service"
	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	r       *gin.Engine
	ctx     context.Context
	service service.AccountService
}

func NewAccountHandler(
	r *gin.Engine,
	ctx context.Context,
	service service.AccountService,
) *accountHandler {
	return &accountHandler{
		r:       r,
		ctx:     ctx,
		service: service,
	}
}

// @Summary Get All Accounts
// @Tags accounts
// @Produce json
// @Success 200 {object}  GetAllAccountsResponse
// @Router /accounts [get]
func (a *accountHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.service.GetAll(ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get One Account By ID
// @Tags accounts
// @Produce json
// @Param id path string true "account ID"
// @Success 200 {object}  GetOneAccountResponse
// @Router /accounts/{id} [get]
func (a *accountHandler) GetOne(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	result, err := a.service.GetOne(ctx, id)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create Account
// @Tags accounts
// @Accept json
// @Produce json
// @Param requestBody body CreateAccountRequest true "Request Body"
// @Success 200 {object} CreateAccountResponse
// @Router /accounts [post]
func (a *accountHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.CreateAccountRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errors.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := a.service.Create(ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// @Summary Update Account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param requestBody body UpdateAccountRequest true "Request Body"
// @Success 200 {object} UpdateByIdAccountResponse
// @Router /accounts [put]
func (a *accountHandler) UpdateById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	payload := dto.UpdateAccountRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errors.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, errData := a.service.UpdateById(ctx, id, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete Account
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 204 {object} DeleteByIdAccountResponse
// @Router /accounts [delete]
func (a *accountHandler) DeleteById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	result, errData := a.service.DeleteById(ctx, id)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusNoContent, result)
}
