package entity

import (
	"time"

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
