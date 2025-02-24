package entity

import (
	"time"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/google/uuid"
)

type Account struct {
	Id            uuid.UUID
	AccountNumber string
	AccountHolder string
	Balance       float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Accounts []Account

func (a *Account) ToAccountResponseDTO() *dto.AccountResponseDTO {
	return &dto.AccountResponseDTO{
		Id:            a.Id,
		AccountNumber: a.AccountNumber,
		AccountHolder: a.AccountHolder,
		Balance:       a.Balance,
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}

func (a Accounts) ToSliceOfAccountsResponseDTO() []dto.AccountResponseDTO {
	result := []dto.AccountResponseDTO{}
	for _, as := range a {
		result = append(result, *as.ToAccountResponseDTO())
	}

	return result
}
