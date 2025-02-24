package account_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
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

func (a *accountPG) GetAll(ctx context.Context) ([]entity.Account, errors.MessageErr) {
	rows, err := a.db.QueryContext(ctx, GET_ALL_ACCOUNTS)

	if err != nil {
		log.Printf("db get all accounts: %s\n", err.Error())
		return nil, errors.NewInternalServerError()
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
			return nil, errors.NewInternalServerError()
		}

		result = append(result, account)
	}

	return result, nil
}

func (a *accountPG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Account, errors.MessageErr) {
	_, err := a.db.QueryContext(ctx, GET_ACCOUNT_BY_ID, id)
	account := entity.Account{}

	if err != nil {
		log.Printf("db get one account by id: %s\n", err.Error())
		return nil, errors.NewInternalServerError()
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
			return nil, errors.NewNotFoundError("account was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &account, nil
}

func (a *accountPG) GetOneByAccountNumber(ctx context.Context, accountNum string) (*entity.Account, errors.MessageErr) {
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
			return nil, errors.NewNotFoundError("account was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &account, nil
}

func (a *accountPG) Create(ctx context.Context, account entity.Account) (*entity.Account, errors.MessageErr) {
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
		return nil, errors.NewInternalServerError()
	}

	return &newAccount, nil
}
func (a *accountPG) UpdateById(ctx context.Context, account entity.Account) (*entity.Account, errors.MessageErr) {
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
			return nil, errors.NewNotFoundError("account was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &response, nil
}
func (a *accountPG) DeleteById(ctx context.Context, id uuid.UUID) errors.MessageErr {
	if _, err := a.db.ExecContext(
		ctx,
		DELETE_ACCOUNT,
		id,
	); err != nil {
		log.Printf("db delete transaction by id: %s\n", err.Error())
		return errors.NewInternalServerError()
	}

	return nil
}
