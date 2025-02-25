package account_repo

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.Account, errs.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.Account, errs.MessageErr)
	Create(ctx context.Context, account entity.Account) (*entity.Account, errs.MessageErr)
	GetOneByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, errs.MessageErr)
	UpdateById(ctx context.Context, account entity.Account) (*entity.Account, errs.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr
	TransferById(ctx context.Context, accountFromId, accountToId entity.Account, amount float64) (*entity.Transaction, errs.MessageErr)
}
