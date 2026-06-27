package usecase

import (
	"testing"

	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/messaging"
	"github.com/joaodddev/transaction-processing-engine/internal/mocks"
)

func TestCreateTransactionUseCase_Execute(
	t *testing.T,
) {

	repository := mocks.NewTransactionRepositoryMock()

	publisher domain.MessagePublisher

	useCase := NewCreateTransactionUseCase(
		repository,
		publisher,
	)

	transaction, err := useCase.Execute(
		"account-123",
		500,
		"BRL",
	)

	if err != nil {
		t.Fatalf(
			"expected nil error, got %v",
			err,
		)
	}

	if transaction == nil {
		t.Fatal("expected transaction, got nil")
	}

	savedTransaction, _ :=
		repository.FindByID(transaction.ID)

	if savedTransaction == nil {
		t.Fatal("transaction was not persisted")
	}
}
