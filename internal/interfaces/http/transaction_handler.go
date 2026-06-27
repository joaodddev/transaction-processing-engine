package http

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
)

type TransactionHandler struct {
	createTransactionUseCase *usecase.CreateTransactionUseCase
}

func NewTransactionHandler(
	createTransactionUseCase *usecase.CreateTransactionUseCase,
) *TransactionHandler {

	return &TransactionHandler{
		createTransactionUseCase: createTransactionUseCase,
	}
}

type CreateTransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
}

func (h *TransactionHandler) CreateTransaction(
	w nethttp.ResponseWriter,
	r *nethttp.Request,
) {

	var request CreateTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusBadRequest)
		return
	}

	output, err := h.createTransactionUseCase.Execute(
		usecase.CreateTransactionInput{
			AccountID: request.AccountID,
			Amount:    request.Amount,
			Currency:  request.Currency,
		},
	)

	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusCreated)

	json.NewEncoder(w).Encode(output)
}
