package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	application "github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
	"github.com/joaodddev/transaction-processing-engine/internal/application/worker"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/database"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/messaging"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/repository"
	handlers "github.com/joaodddev/transaction-processing-engine/internal/interfaces/http"
	customMiddleware "github.com/joaodddev/transaction-processing-engine/internal/interfaces/http/middleware"
	"github.com/joaodddev/transaction-processing-engine/internal/observability/metrics"
)

func main() {

	metrics.Register()

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

	healthHandler :=
		handlers.NewHealthHandler()

	router := chi.NewRouter()

	router.Use(customMiddleware.Logging)

	router.Post(
		"/transactions",
		transactionHandler.CreateTransaction,
	)

	router.Get(
		"/transactions/{id}",
		transactionHandler.GetTransaction,
	)

	router.Get(
		"/health",
		healthHandler.HealthCheck,
	)

	router.Handle(
		"/metrics",
		promhttp.Handler(),
	)

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
