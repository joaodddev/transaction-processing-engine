package worker

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/joaodddev/transaction-processing-engine/internal/application/queue"
	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

type TransactionWorker struct {
	id         int
	queue      *queue.TransactionQueue
	repository domain.TransactionRepository
}

func NewTransactionWorker(
	id int,
	queue *queue.TransactionQueue,
	repository domain.TransactionRepository,
) *TransactionWorker {

	return &TransactionWorker{
		id:         id,
		queue:      queue,
		repository: repository,
	}
}

func (w *TransactionWorker) Start() {

	go func() {

		for transactionID := range w.queue.Consume() {

			log.Printf(
				"[worker-%d] processing transaction %s",
				w.id,
				transactionID.String(),
			)

			w.process(transactionID)
		}
	}()
}

func (w *TransactionWorker) process(id uuid.UUID) {

	transaction, err := w.repository.FindByID(id)

	if err != nil || transaction == nil {
		return
	}

	transaction.StartProcessing()

	time.Sleep(3 * time.Second)

	if rand.Intn(100) < 80 {
		transaction.Approve()
	} else {
		transaction.Reject()
	}

	log.Printf(
		"[worker-%d] transaction %s finished with status %s",
		w.id,
		transaction.ID.String(),
		transaction.Status,
	)
}
