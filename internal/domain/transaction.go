package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	StatusPending    TransactionStatus = "PENDING"
	StatusProcessing TransactionStatus = "PROCESSING"
	StatusApproved   TransactionStatus = "APPROVED"
	StatusRejected   TransactionStatus = "REJECTED"
	StatusFailed     TransactionStatus = "FAILED"
)

type Transaction struct {
	ID          uuid.UUID
	AccountID   string
	Amount      float64
	Currency    string
	Status      TransactionStatus
	CreatedAt   time.Time
	ProcessedAt *time.Time
}

func NewTransaction(
	accountID string,
	amount float64,
	currency string,
) (*Transaction, error) {

	if accountID == "" {
		return nil, errors.New("account id is required")
	}

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	return &Transaction{
		ID:        uuid.New(),
		AccountID: accountID,
		Amount:    amount,
		Currency:  currency,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}, nil
}
