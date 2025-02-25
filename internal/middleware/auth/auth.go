package auth

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/config"
	"github.com/adityatresnobudi/go-restapi-gin/internal/domain/user/service"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/internal_jwt"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authentication() gin.HandlerFunc
}

type authMiddlewareIMPL struct {
	ctx         context.Context
	internalJwt internal_jwt.InternalJwt
	cfg         config.Config
	userService service.UserService
}

func NewAuthMiddleware(
	ctx context.Context,
	internalJwt internal_jwt.InternalJwt,
	cfg config.Config,
	userService service.UserService,
) AuthMiddleware {
	return &authMiddlewareIMPL{
		ctx:         ctx,
		internalJwt: internalJwt,
		cfg:         cfg,
		userService: userService,
	}
}

func (a *authMiddlewareIMPL) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		mapClaims, err := a.internalJwt.ValidateBearerToken(token, a.cfg.Jwt.SecretKey)

		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		id, ok := mapClaims["id"].(float64)

		if !ok {
			err := errs.NewUnauthenticatedError("invalid token.")
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		user, err := a.userService.GetById(a.ctx, int(id))
		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		c.Set("userId", int(id))
		c.Set("role", user.Roles)

		c.Next()
	}
}
