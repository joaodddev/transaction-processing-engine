package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	application "github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
	"github.com/joaodddev/transaction-processing-engine/internal/application/worker"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/database"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/messaging"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/repository"
	handlers "github.com/joaodddev/transaction-processing-engine/internal/interfaces/http"
)

func main() {

	db, err := database.NewPostgresConnection()

	if err != nil {
		panic(err)
	}

	repository :=
		repository.NewPostgresTransactionRepository(db)

	rabbitMQ, err := messaging.NewRabbitMQ()

	if err != nil {
		panic(err)
	}

	defer rabbitMQ.Connection.Close()
	defer rabbitMQ.Channel.Close()

	publisher :=
		messaging.NewPublisher(rabbitMQ.Channel)

	createTransactionUseCase :=
		application.NewCreateTransactionUseCase(
			repository,
			publisher,
		)

	if err := worker.StartConsumer(
		rabbitMQ.Channel,
		repository,
	); err != nil {

		panic(err)
	}

	transactionHandler :=
		handlers.NewTransactionHandler(
			createTransactionUseCase,
			repository,
			nil,
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
