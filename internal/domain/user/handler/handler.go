package handler

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/domain/user/service"
	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
	r       *gin.Engine
	ctx     context.Context
}

func NewUserHandler(service service.UserService, r *gin.Engine, ctx context.Context) *userHandler {
	return &userHandler{
		service: service,
		r:       r,
		ctx:     ctx,
	}
}

func (u *userHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.RegisterRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
	}

	result, errData := u.service.Create(ctx, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
	}

	c.JSON(http.StatusCreated, result)
}

func (u *userHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.LoginRequestDTO{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
	}

	result, errData := u.service.Login(ctx, payload)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
	}

	c.JSON(http.StatusOK, result)
}
