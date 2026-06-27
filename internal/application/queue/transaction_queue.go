package queue

import "github.com/google/uuid"

type TransactionQueue struct {
	channel chan uuid.UUID
}

func NewTransactionQueue(bufferSize int) *TransactionQueue {
	return &TransactionQueue{
		channel: make(chan uuid.UUID, bufferSize),
	}
}

func (q *TransactionQueue) Publish(id uuid.UUID) {
	q.channel <- id
}

func (q *TransactionQueue) Consume() <-chan uuid.UUID {
	return q.channel
}
