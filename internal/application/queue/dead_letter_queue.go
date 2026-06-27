package queue

import (
	"sync"

	"github.com/google/uuid"
)

type DeadLetterQueue struct {
	mu             sync.RWMutex
	transactionsID []uuid.UUID
}

func NewDeadLetterQueue() *DeadLetterQueue {
	return &DeadLetterQueue{
		transactionsID: make([]uuid.UUID, 0),
	}
}

func (q *DeadLetterQueue) Add(id uuid.UUID) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.transactionsID = append(q.transactionsID, id)
}

func (q *DeadLetterQueue) GetAll() []uuid.UUID {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.transactionsID
}
