package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

type InMemoryTransactionRepository struct {
	mu           sync.RWMutex
	transactions map[uuid.UUID]*domain.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[uuid.UUID]*domain.Transaction),
	}
}

func (r *InMemoryTransactionRepository) Save(
	transaction *domain.Transaction,
) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.transactions[transaction.ID] = transaction

	return nil
}

func (r *InMemoryTransactionRepository) FindByID(
	id uuid.UUID,
) (*domain.Transaction, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	transaction, exists := r.transactions[id]

	if !exists {
		return nil, nil
	}

	return transaction, nil
}

func (r *InMemoryTransactionRepository) FindAll() ([]*domain.Transaction, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	transactions := make([]*domain.Transaction, 0)

	for _, transaction := range r.transactions {
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
