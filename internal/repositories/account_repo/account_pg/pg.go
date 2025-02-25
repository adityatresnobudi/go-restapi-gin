package account_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/google/uuid"
)

type accountPG struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) account_repo.Repository {
	return &accountPG{
		db: db,
	}
}

func (a *accountPG) GetAll(ctx context.Context) ([]entity.Account, errs.MessageErr) {
	rows, err := a.db.QueryContext(ctx, GET_ALL_ACCOUNTS)

	if err != nil {
		log.Printf("db get all accounts: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	result := []entity.Account{}

	for rows.Next() {
		account := entity.Account{}

		if err = rows.Scan(
			&account.Id,
			&account.AccountNumber,
			&account.AccountHolder,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		); err != nil {
			log.Printf("db scan get all accounts: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		result = append(result, account)
	}

	return result, nil
}

func (a *accountPG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Account, errs.MessageErr) {
	_, err := a.db.QueryContext(ctx, GET_ACCOUNT_BY_ID, id)
	account := entity.Account{}

	if err != nil {
		log.Printf("db get one account by id: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	if err := a.db.QueryRowContext(
		ctx,
		GET_ACCOUNT_BY_ID,
		id,
	).Scan(
		&account.Id,
		&account.AccountNumber,
		&account.AccountHolder,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one account by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &account, nil
}

func (a *accountPG) GetOneByAccountNumber(ctx context.Context, accountNum string) (*entity.Account, errs.MessageErr) {
	account := entity.Account{}

	if err := a.db.QueryRowContext(
		ctx,
		GET_ACCOUNT_BY_ACCOUNTNUM,
		accountNum,
	).Scan(
		&account.Id,
		&account.AccountNumber,
		&account.AccountHolder,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one account by accountNum: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &account, nil
}

func (a *accountPG) Create(ctx context.Context, account entity.Account) (*entity.Account, errs.MessageErr) {
	newAccount := entity.Account{}

	if err := a.db.QueryRowContext(
		ctx,
		INSERT_ACCOUNT,
		account.AccountNumber,
		account.AccountHolder,
		account.Balance,
	).Scan(
		&newAccount.Id,
		&newAccount.AccountNumber,
		&newAccount.AccountHolder,
		&newAccount.Balance,
		&newAccount.CreatedAt,
		&newAccount.UpdatedAt,
	); err != nil {
		log.Printf("db scan create account: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	return &newAccount, nil
}
func (a *accountPG) UpdateById(ctx context.Context, account entity.Account) (*entity.Account, errs.MessageErr) {
	response := entity.Account{}

	if err := a.db.QueryRowContext(
		ctx,
		UPDATE_ACCOUNT,
		account.AccountHolder,
		account.Balance,
		account.Id,
	).Scan(
		&response.Id,
		&response.AccountNumber,
		&response.AccountHolder,
		&response.Balance,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		log.Printf("db scan update account by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account was not found")
		}
		return nil, errs.NewInternalServerError()
	}

	return &response, nil
}
func (a *accountPG) DeleteById(ctx context.Context, id uuid.UUID) errs.MessageErr {
	if _, err := a.db.ExecContext(
		ctx,
		DELETE_ACCOUNT,
		id,
	); err != nil {
		log.Printf("db delete transaction by id: %s\n", err.Error())
		return errs.NewInternalServerError()
	}

	return nil
}
func (a *accountPG) TransferById(ctx context.Context, accountFromId, accountToId entity.Account, amount float64) (*entity.Transaction, errs.MessageErr) {
	newTransaction := entity.Transaction{}
	tx, err := a.db.BeginTx(ctx, nil)
    if err != nil {
        log.Printf("tx begin: %s\n", err.Error())
    }
	defer tx.Rollback()


	if _, err := tx.ExecContext(
		ctx,
		UPDATE_BALANCE,
		accountFromId.Balance,
		accountFromId.Id,
	); err != nil {
		log.Printf("tx update transfer from by id: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	if _, err := tx.ExecContext(
		ctx,
		UPDATE_BALANCE,
		accountToId.Balance,
		accountToId.Id,
	); err != nil {
		log.Printf("tx update transfer to by id: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	if err := tx.QueryRowContext(
		ctx, 
		INSERT_TRANSACTION,
		accountFromId.Id,
		accountToId.Id,
		amount,
	).Scan(
		&newTransaction.Id,
		&newTransaction.AccountIdFrom,
		&newTransaction.AccountIdTo,
		&newTransaction.Amount,
		&newTransaction.CreatedAt,
		&newTransaction.UpdatedAt,
	); err != nil {
		log.Printf("tx create transaction: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	if err = tx.Commit(); err != nil {
		log.Printf("tx commit update transfer: %s\n", err.Error())
        return nil, errs.NewInternalServerError()
    }

	return &newTransaction, nil
}