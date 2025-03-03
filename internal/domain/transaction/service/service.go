package service

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/transaction_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errs"
	"github.com/google/uuid"
)

type TransactionService interface {
	GetTransactionById(ctx context.Context, id string) (*dto.GetTransactionByIdResponseDTO, errs.MessageErr)
	Create(ctx context.Context, transaction dto.CreateTransactionRequestDTO) (*dto.CreateTransactionResponseDTO, errs.MessageErr)
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

func (t *transactionServiceIMPL) GetTransactionById(ctx context.Context, id string) (*dto.GetTransactionByIdResponseDTO, errs.MessageErr) {
	username := ctx.Value("username")
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errs.NewBadRequest("id has to be a valid uuid")
	}

	user, err := t.accountRepo.GetOneByUsername(ctx, username.(string))
	if err != nil {
		return nil, err
	}
	account, err := t.accountRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	if user.AccountNumber != account.AccountNumber {
		return nil, errs.NewBadRequest("id has to be your id")
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
) (*dto.CreateTransactionResponseDTO, errs.MessageErr) {
	username := ctx.Value("username")
	parseFromId, _ := uuid.Parse(transaction.FromAccountId)
	parseToId, _ := uuid.Parse(transaction.ToAccountId)

	user, err := t.accountRepo.GetOneByUsername(ctx, username.(string))
	if err != nil {
		return nil, err
	}

	existingAccountFrom, err := t.accountRepo.GetOneById(
		ctx,
		parseFromId,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if user.AccountNumber != existingAccountFrom.AccountNumber {
		return nil, errs.NewBadRequest("cannot transfer from other account")
	}

	if existingAccountFrom.Balance < transaction.Amount {
		return nil, errs.NewBadRequest("balance not enough")
	}

	existingAccountTo, err := t.accountRepo.GetOneById(
		ctx,
		parseToId,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if parseFromId == parseToId {
		return nil, errs.NewBadRequest("cannot transfer to the same account")
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

	resp, err := t.accountRepo.TransferById(ctx, accountFrom, accountTo, transaction.Amount)
	if err != nil {
		return nil, err
	}

	result := dto.CreateTransactionResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Transaction created successfully"},
		Data:                  *resp.ToTransactionResponseDTO(),
	}

	return &result, nil
}
