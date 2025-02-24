package service

import (
	"context"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/adityatresnobudi/go-restapi-gin/internal/entity"
	"github.com/adityatresnobudi/go-restapi-gin/internal/repositories/account_repo"
	"github.com/adityatresnobudi/go-restapi-gin/pkg/errors"
	"github.com/google/uuid"
)

type AccountService interface {
	GetAll(ctx context.Context) (*dto.GetAllAccountsResponseDTO, errors.MessageErr)
	GetOne(ctx context.Context, id string) (*dto.GetOneAccountResponseDTO, errors.MessageErr)
	Create(ctx context.Context, payload dto.CreateAccountRequestDTO) (*dto.CreateAccountResponseDTO, errors.MessageErr)
	UpdateById(ctx context.Context, id string, payload dto.UpdateAccountRequestDTO) (*dto.UpdateAccountResponseDTO, errors.MessageErr)
	DeleteById(ctx context.Context, id string) (*dto.DeleteByIdAccountResponseDTO, errors.MessageErr)
}

type accountServiceIMPL struct {
	accountRepo account_repo.Repository
}

func NewAccountService(accountRepo account_repo.Repository) AccountService {
	return &accountServiceIMPL{
		accountRepo: accountRepo,
	}
}

func (a *accountServiceIMPL) GetAll(ctx context.Context) (*dto.GetAllAccountsResponseDTO, errors.MessageErr) {
	accounts, err := a.accountRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetAllAccountsResponseDTO{
		Data: entity.Accounts(accounts).ToSliceOfAccountsResponseDTO(),
	}

	return &result, nil
}

func (a *accountServiceIMPL) GetOne(ctx context.Context, id string) (*dto.GetOneAccountResponseDTO, errors.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	account, err := a.accountRepo.GetOneById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	result := dto.GetOneAccountResponseDTO{
		Data: *account.ToAccountResponseDTO(),
	}

	return &result, nil
}

func (a *accountServiceIMPL) Create(
	ctx context.Context,
	payload dto.CreateAccountRequestDTO,
) (*dto.CreateAccountResponseDTO, errors.MessageErr) {
	if err := a.createValidator(payload); err != nil {
		return nil, err
	}

	existingAccount, err := a.accountRepo.GetOneByAccountNumber(
		ctx,
		payload.AccountNumber,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, errors.NewConflictError("please choose another account number")
	}

	account := entity.Account{
		AccountNumber: payload.AccountNumber,
		AccountHolder: payload.AccountHolder,
		Balance:       payload.Balance,
	}

	newAccount, err := a.accountRepo.Create(ctx, account)

	if err != nil {
		return nil, err
	}

	result := dto.CreateAccountResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Account created successfully"},
		Data:                  *newAccount.ToCreateAccountResponseDTO(),
	}

	return &result, nil
}

func (a *accountServiceIMPL) UpdateById(ctx context.Context, id string, payload dto.UpdateAccountRequestDTO) (*dto.UpdateAccountResponseDTO, errors.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	if err := a.updateValidator(payload); err != nil {
		return nil, err
	}

	account, err := a.accountRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	account.Id = parsedId
	account.AccountHolder = payload.AccountHolder
	account.Balance = payload.Balance

	response, err := a.accountRepo.UpdateById(ctx, *account)
	if err != nil {
		return nil, err
	}

	result := dto.UpdateAccountResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Account updated successfully"},
		Data:                  *response.ToAccountResponseDTO(),
	}

	return &result, nil
}

func (a *accountServiceIMPL) DeleteById(ctx context.Context, id string) (*dto.DeleteByIdAccountResponseDTO, errors.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	_, err := a.accountRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	err = a.accountRepo.DeleteById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	result := dto.DeleteByIdAccountResponseDTO{
		CommonBaseResponseDTO: dto.CommonBaseResponseDTO{Message: "Account deleted successfully"},
	}

	return &result, nil
}
