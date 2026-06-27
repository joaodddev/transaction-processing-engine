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
	ID          uuid.UUID         `json:"id"`
	AccountID   string            `json:"account_id"`
	Amount      float64           `json:"amount"`
	Currency    string            `json:"currency"`
	Status      TransactionStatus `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	ProcessedAt *time.Time        `json:"processed_at,omitempty"`
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

	if currency == "" {
		return nil, errors.New("currency is required")
	}

	return &Transaction{
		ID:          uuid.New(),
		AccountID:   accountID,
		Amount:      amount,
		Currency:    currency,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		ProcessedAt: nil,
	}, nil
}

func (t *Transaction) StartProcessing() {
	t.Status = StatusProcessing
}

func (t *Transaction) Approve() {
	now := time.Now()

	t.Status = StatusApproved
	t.ProcessedAt = &now
}

func (t *Transaction) Reject() {
	now := time.Now()

	t.Status = StatusRejected
	t.ProcessedAt = &now
}

func (t *Transaction) Fail() {
	now := time.Now()

	t.Status = StatusFailed
	t.ProcessedAt = &now
}
