package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	application "github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/repository"
	handlers "github.com/joaodddev/transaction-processing-engine/internal/interfaces/http"
)

func main() {

	repository := repository.NewInMemoryTransactionRepository()

	createTransactionUseCase :=
		application.NewCreateTransactionUseCase(repository)

	transactionHandler :=
		handlers.NewTransactionHandler(
			createTransactionUseCase,
		)

	router := chi.NewRouter()

	router.Post(
		"/transactions",
		transactionHandler.CreateTransaction,
	)

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
