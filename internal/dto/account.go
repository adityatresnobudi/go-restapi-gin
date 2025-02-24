package dto

import (
	"time"

	"github.com/google/uuid"
)

type AccountResponseDTO struct {
	Id            uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	AccountNumber string    `json:"account_number" example:"233455011"`
	AccountHolder string    `json:"account_holder" example:"adit"`
	Balance       float64   `json:"balance" example:"10.3"`
	CreatedAt     time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
	UpdatedAt     time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
} // @name AccountResponse

type CreateAccountRequestDTO struct {
	AccountNumber string  `json:"account_number" example:"1234567890"`
	AccountHolder string  `json:"account_holder" example:"thomas"`
	Balance       float64 `json:"balance" example:"50000"`
} // @name CreateAccountRequest

type UpdateAccountRequestDTO struct {
	AccountHolder string  `json:"account_holder" example:"thomas"`
	Balance       float64 `json:"balance" example:"50000"`
} // @name UpdateAccountRequest

type GetAllAccountsResponseDTO struct {
	Data []AccountResponseDTO
} // @name GetAllAccountsResponse

type GetOneAccountResponseDTO struct {
	Data AccountResponseDTO `json:"data"`
} // @name GetOneAccountResponse

type CreateAccountResponseDTO struct {
	CommonBaseResponseDTO
	Data AccountResponseDTO `json:"data"`
} // @name CreateAccountResponse

type UpdateAccountResponseDTO struct {
	CommonBaseResponseDTO
	Data AccountResponseDTO `json:"data"`
} // @name UpdateAccountResponse

type DeleteByIdAccountResponseDTO struct {
	CommonBaseResponseDTO
} // @name DeleteAccountResponse
