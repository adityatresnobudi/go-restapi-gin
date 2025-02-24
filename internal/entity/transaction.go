package entity

import (
	"time"

	"github.com/adityatresnobudi/go-restapi-gin/internal/dto"
	"github.com/google/uuid"
)

type Transaction struct {
	Id            uuid.UUID
	AccountIdFrom string
	AccountIdTo   string
	Amount        float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Transactions []Transaction

func (t *Transaction) ToTransactionResponseDTO() *dto.TransactionResponseDTO {
	return &dto.TransactionResponseDTO{
		Id:            t.Id,
		AccountIdFrom: t.AccountIdFrom,
		AccountIdTo:   t.AccountIdTo,
		Amount:        t.Amount,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

func (t Transactions) ToSliceOfTransactionsResponseDTO() []dto.TransactionResponseDTO {
	result := []dto.TransactionResponseDTO{}
	for _, tx := range t {
		result = append(result, *tx.ToTransactionResponseDTO())
	}

	return result
}
