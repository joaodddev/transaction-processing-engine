package http

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/joaodddev/transaction-processing-engine/internal/application/usecase"
	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

type TransactionHandler struct {
	createTransactionUseCase *usecase.CreateTransactionUseCase
	repository               domain.TransactionRepository
}

func NewTransactionHandler(
	createTransactionUseCase *usecase.CreateTransactionUseCase,
	repository domain.TransactionRepository,
) *TransactionHandler {

	return &TransactionHandler{
		createTransactionUseCase: createTransactionUseCase,
		repository:               repository,
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

func (h *TransactionHandler) GetTransaction(
	w nethttp.ResponseWriter,
	r *nethttp.Request,
) {

	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		nethttp.Error(w, "invalid id", nethttp.StatusBadRequest)
		return
	}

	transaction, err := h.repository.FindByID(id)

	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)
		return
	}

	if transaction == nil {
		nethttp.Error(w, "transaction not found", nethttp.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transaction)
}
