package account_repo

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.Account, errors.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.Account, errors.MessageErr)
	Create(ctx context.Context, account entity.Account) (*entity.Account, errors.MessageErr)
	GetOneByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, errors.MessageErr)
	UpdateById(ctx context.Context, account entity.Account) (*entity.Account, errors.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errors.MessageErr
	TransferById(ctx context.Context, accountFromId, accountToId entity.Account) (errors.MessageErr)
}
