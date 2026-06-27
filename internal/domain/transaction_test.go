package domain

import "testing"

func TestNewTransaction_ShouldCreateTransactionSuccessfully(
	t *testing.T,
) {

	transaction, err := NewTransaction(
		"account-123",
		100.50,
		"BRL",
	)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if transaction.AccountID != "account-123" {
		t.Errorf(
			"expected account id account-123, got %s",
			transaction.AccountID,
		)
	}

	if transaction.Amount != 100.50 {
		t.Errorf(
			"expected amount 100.50, got %f",
			transaction.Amount,
		)
	}

	if transaction.Currency != "BRL" {
		t.Errorf(
			"expected currency BRL, got %s",
			transaction.Currency,
		)
	}

	if transaction.Status != StatusPending {
		t.Errorf(
			"expected status PENDING, got %s",
			transaction.Status,
		)
	}
}

func TestNewTransaction_ShouldReturnErrorForInvalidAmount(
	t *testing.T,
) {

	_, err := NewTransaction(
		"account-123",
		0,
		"BRL",
	)

	if err == nil {
		t.Error("expected error, got nil")
	}
}
