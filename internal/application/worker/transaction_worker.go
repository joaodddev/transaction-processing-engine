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
	dlq        *queue.DeadLetterQueue
	repository domain.TransactionRepository
}

func NewTransactionWorker(
	id int,
	queue *queue.TransactionQueue,
	dlq *queue.DeadLetterQueue,
	repository domain.TransactionRepository,
) *TransactionWorker {

	return &TransactionWorker{
		id:         id,
		queue:      queue,
		dlq:        dlq,
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

	success := rand.Intn(100) < 50

	if success {
		transaction.Approve()

		log.Printf(
			"[worker-%d] transaction %s approved",
			w.id,
			transaction.ID,
		)

		return
	}

	transaction.IncrementRetry()

	log.Printf(
		"[worker-%d] transaction %s failed - retry %d/%d",
		w.id,
		transaction.ID,
		transaction.RetryCount,
		transaction.MaxRetries,
	)

	if transaction.CanRetry() {

		transaction.Status = domain.StatusPending

		go func() {
			time.Sleep(2 * time.Second)
			w.queue.Publish(transaction.ID)
		}()

		return
	}

	transaction.Fail()

	w.dlq.Add(transaction.ID)

	log.Printf(
		"[worker-%d] transaction %s moved to DLQ",
		w.id,
		transaction.ID,
	)
}
