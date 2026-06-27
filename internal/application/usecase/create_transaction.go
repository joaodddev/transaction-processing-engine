package usecase

import (
	"encoding/json"

	"github.com/joaodddev/transaction-processing-engine/internal/domain"
	"github.com/joaodddev/transaction-processing-engine/internal/infrastructure/messaging"
)

type CreateTransactionInput struct {
	AccountID string
	Amount    float64
	Currency  string
}

type CreateTransactionOutput struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

type CreateTransactionUseCase struct {
	repository domain.TransactionRepository
	publisher  *messaging.Publisher
}

func NewCreateTransactionUseCase(
	repository domain.TransactionRepository,
	publisher *messaging.Publisher,
) *CreateTransactionUseCase {

	return &CreateTransactionUseCase{
		repository: repository,
		publisher:  publisher,
	}
}

func (uc *CreateTransactionUseCase) Execute(
	input CreateTransactionInput,
) (*CreateTransactionOutput, error) {

	transaction, err := domain.NewTransaction(
		input.AccountID,
		input.Amount,
		input.Currency,
	)

	if err != nil {
		return nil, err
	}

	if err := uc.repository.Save(transaction); err != nil {
		return nil, err
	}

	payload, _ := json.Marshal(transaction)

	if err := uc.publisher.Publish(payload); err != nil {
		return nil, err
	}

	return &CreateTransactionOutput{
		ID:        transaction.ID.String(),
		AccountID: transaction.AccountID,
		Amount:    transaction.Amount,
		Currency:  transaction.Currency,
		Status:    string(transaction.Status),
		CreatedAt: transaction.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
