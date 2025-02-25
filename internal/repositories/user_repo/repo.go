package user_repo

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
)

type Repository interface {
	Create(ctx context.Context, user entity.User) errs.MessageErr
	GetByUsername(ctx context.Context, username string) (*entity.User, errs.MessageErr)
	GetById(ctx context.Context, id int) (*entity.User, errs.MessageErr)
}
