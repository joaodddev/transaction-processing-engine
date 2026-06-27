package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	applicationQueue "github.com/joaodddev/transaction-processing-engine/internal/application/queue"
	application "github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
	"github.com/joaodddev/transaction-processing-engine/internal/application/worker"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/repository"
	handlers "github.com/joaodddev/transaction-processing-engine/internal/interfaces/http"
)

func main() {

	repository := repository.NewInMemoryTransactionRepository()

	transactionQueue := applicationQueue.NewTransactionQueue(100)

	createTransactionUseCase :=
		application.NewCreateTransactionUseCase(
			repository,
			transactionQueue,
		)

	for i := 1; i <= 3; i++ {
		worker := worker.NewTransactionWorker(
			i,
			transactionQueue,
			repository,
		)

		worker.Start()
	}

	transactionHandler :=
		handlers.NewTransactionHandler(
			createTransactionUseCase,
			repository,
		)

	router := chi.NewRouter()

	router.Post(
		"/transactions",
		transactionHandler.CreateTransaction,
	)

	router.Get(
		"/transactions/{id}",
		transactionHandler.GetTransaction,
	)

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
