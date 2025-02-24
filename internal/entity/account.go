package entity

import (
	"time"

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
