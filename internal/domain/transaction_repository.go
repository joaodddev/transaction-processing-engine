package domain

import "github.com/google/uuid"

type TransactionRepository interface {
	Save(transaction *Transaction) error
	FindByID(id uuid.UUID) (*Transaction, error)
	FindAll() ([]*Transaction, error)
}
