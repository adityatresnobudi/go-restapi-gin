package transaction_repo

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errors.MessageErr)
	GetTransactionById(ctx context.Context, id string) ([]entity.Transaction, errors.MessageErr)
}
