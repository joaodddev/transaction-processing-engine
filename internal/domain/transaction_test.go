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

func TestTransaction_Approve(t *testing.T) {

	transaction, _ := NewTransaction(
		"account-123",
		100,
		"BRL",
	)

	transaction.Approve()

	if transaction.Status != StatusApproved {
		t.Errorf(
			"expected APPROVED, got %s",
			transaction.Status,
		)
	}

	if transaction.ProcessedAt == nil {
		t.Error("expected processed_at to be filled")
	}
}

func TestTransaction_Fail(t *testing.T) {

	transaction, _ := NewTransaction(
		"account-123",
		100,
		"BRL",
	)

	transaction.Fail()

	if transaction.Status != StatusFailed {
		t.Errorf(
			"expected FAILED, got %s",
			transaction.Status,
		)
	}
}

func TestTransaction_Retry(t *testing.T) {

	transaction, _ := NewTransaction(
		"account-123",
		100,
		"BRL",
	)

	transaction.IncrementRetry()

	if transaction.RetryCount != 1 {
		t.Errorf(
			"expected retry count 1, got %d",
			transaction.RetryCount,
		)
	}
}
