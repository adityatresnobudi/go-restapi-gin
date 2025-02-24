package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionResponseDTO struct {
	Id            uuid.UUID `json:"id" example:"ddcba067-37a7-4381-8f4a-c2c0bc2891e0"`
	AccountIdFrom string    `json:"from_account_id" example:"233455011"`
	AccountIdTo   string    `json:"to_account_id" example:"233455011"`
	Amount        float64   `json:"amount" example:"10.3"`
	CreatedAt     time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt     time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
}

type CreateTransactionRequestDTO struct {
	FromAccountId string  `json:"from_account_id"`
	ToAccountId   string  `json:"to_account_id"`
	Amount        float64 `json:"amount"`
}

type GetTransactionByIdResponseDTO struct {
	CommonBaseResponseDTO
	Data []TransactionResponseDTO `json:"data"`
}

type CreateTransactionResponseDTO struct {
	CommonBaseResponseDTO
	Data TransactionResponseDTO `json:"data"`
}
