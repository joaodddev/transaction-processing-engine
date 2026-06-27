package mocks

import (
	"github.com/google/uuid"

	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

type TransactionRepositoryMock struct {
	Transactions map[uuid.UUID]*domain.Transaction
}

func NewTransactionRepositoryMock() *TransactionRepositoryMock {

	return &TransactionRepositoryMock{
		Transactions: make(map[uuid.UUID]*domain.Transaction),
	}
}

func (m *TransactionRepositoryMock) Save(
	transaction *domain.Transaction,
) error {

	m.Transactions[transaction.ID] = transaction

	return nil
}

func (m *TransactionRepositoryMock) FindByID(
	id uuid.UUID,
) (*domain.Transaction, error) {

	transaction, exists := m.Transactions[id]

	if !exists {
		return nil, nil
	}

	return transaction, nil
}

func (m *TransactionRepositoryMock) FindAll() ([]*domain.Transaction, error) {

	transactions := make([]*domain.Transaction, 0)

	for _, transaction := range m.Transactions {
		transactions = append(
			transactions,
			transaction,
		)
	}

	return transactions, nil
}
