package service

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/transaction_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
	"github.com/google/uuid"
)

type TransactionService interface {
	GetTransactionById(ctx context.Context, id string) (*dto.GetTransactionByIdResponseDTO, errors.MessageErr)
	Create(ctx context.Context, transaction dto.CreateTransactionRequestDTO) (*dto.CreateTransactionResponseDTO, errors.MessageErr)
}

type transactionServiceIMPL struct {
	transactionRepo transaction_repo.Repository
	accountRepo     account_repo.Repository
}

func NewTransactionService(transactionRepo transaction_repo.Repository, accountRepo account_repo.Repository) TransactionService {
	return &transactionServiceIMPL{
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}

func (t *transactionServiceIMPL) GetTransactionById(ctx context.Context, id string) (*dto.GetTransactionByIdResponseDTO, errors.MessageErr) {
	_, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	transactions, err := t.transactionRepo.GetTransactionById(ctx, id)

	if err != nil {
		return nil, err
	}

	result := dto.GetTransactionByIdResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "OK"},
		Data:                  entity.Transactions(transactions).ToSliceOfTransactionsResponseDTO(),
	}

	return &result, nil
}

func (t *transactionServiceIMPL) Create(
	ctx context.Context,
	transaction dto.CreateTransactionRequestDTO,
) (*dto.CreateTransactionResponseDTO, errors.MessageErr) {
	parseFromId, _ := uuid.Parse(transaction.FromAccountId)
	parseToId, _ := uuid.Parse(transaction.ToAccountId)
	existingAccountFrom, err := t.accountRepo.GetOneById(
		ctx,
		parseFromId,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if existingAccountFrom.Balance < transaction.Amount {
		return nil, errors.NewBadRequest("balance not enough")
	}

	existingAccountTo, err := t.accountRepo.GetOneById(
		ctx,
		parseToId,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	newTransaction := entity.Transaction{
		Id:            parseFromId,
		AccountIdFrom: transaction.FromAccountId,
		AccountIdTo:   transaction.ToAccountId,
		Amount:        transaction.Amount,
	}

	accountFrom := entity.Account{
		Id:            parseFromId,
		AccountHolder: existingAccountFrom.AccountHolder,
		Balance:       existingAccountFrom.Balance - transaction.Amount,
	}

	accountTo := entity.Account{
		Id:            parseToId,
		AccountHolder: existingAccountTo.AccountHolder,
		Balance:       existingAccountTo.Balance + transaction.Amount,
	}

	resp, err := t.transactionRepo.Create(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	_, err = t.accountRepo.UpdateById(ctx, accountFrom)
	if err != nil {
		return nil, err
	}

	_, err = t.accountRepo.UpdateById(ctx, accountTo)
	if err != nil {
		return nil, err
	}

	result := dto.CreateTransactionResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Transaction created successfully"},
		Data:                  *resp.ToTransactionResponseDTO(),
	}

	return &result, nil
}
