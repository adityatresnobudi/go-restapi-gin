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
	ctx     context.Context
	r       *gin.Engine
	service service.UserService
}

func NewUserHandler(ctx context.Context,r *gin.Engine, service service.UserService) *userHandler {
	return &userHandler{
		ctx:     ctx,
		r:       r,
		service: service,
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
