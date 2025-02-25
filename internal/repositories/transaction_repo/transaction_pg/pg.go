package transaction_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/transaction_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
)

type transactionPG struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) transaction_repo.Repository {
	return &transactionPG{
		db: db,
	}
}

func (a *transactionPG) Create(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errs.MessageErr) {
	newTransction := entity.Transaction{}

	if err := a.db.QueryRowContext(
		ctx,
		INSERT_TRANSACTION,
		transaction.AccountIdFrom,
		transaction.AccountIdTo,
		transaction.Amount,
	).Scan(
		&newTransction.Id,
		&newTransction.AccountIdFrom,
		&newTransction.AccountIdTo,
		&newTransction.Amount,
		&newTransction.CreatedAt,
		&newTransction.UpdatedAt,
	); err != nil {
		log.Printf("db scan create account: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	return &newTransction, nil
}

func (t *transactionPG) GetTransactionById(ctx context.Context, id string) ([]entity.Transaction, errs.MessageErr) {
	rows, err := t.db.QueryContext(ctx, GET_ALL_TRANSACTIONS_BY_ID, id, id)

	if err != nil {
		log.Printf("db get all transactions by id: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	result := []entity.Transaction{}

	for rows.Next() {
		transaction := entity.Transaction{}

		if err = rows.Scan(
			&transaction.Id,
			&transaction.AccountIdFrom,
			&transaction.AccountIdTo,
			&transaction.Amount,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		); err != nil {
			log.Printf("db scan get all transactions by id: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		result = append(result, transaction)
	}

	return result, nil
}
