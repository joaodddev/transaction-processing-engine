package domain

import "testing"

func TestShouldCreateTransaction(t *testing.T) {

	transaction, err := NewTransaction(
		"acc-123",
		100,
		"BRL",
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if transaction.Status != StatusPending {
		t.Errorf("expected PENDING status")
	}
}

func TestShouldFailWhenAmountIsInvalid(t *testing.T) {

	_, err := NewTransaction(
		"acc-123",
		0,
		"BRL",
	)

	if err == nil {
		t.Error("expected validation error")
	}
}
