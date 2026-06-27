package repository

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/joaodddev/transaction-processing-engine/internal/domain"
)

type PostgresTransactionRepository struct {
	db *sql.DB
}

func NewPostgresTransactionRepository(
	db *sql.DB,
) *PostgresTransactionRepository {

	return &PostgresTransactionRepository{
		db: db,
	}
}

func (r *PostgresTransactionRepository) Save(
	transaction *domain.Transaction,
) error {

	query := `
	INSERT INTO transactions (
		id,
		account_id,
		amount,
		currency,
		status,
		retry_count,
		max_retries,
		created_at,
		processed_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	ON CONFLICT(id)
	DO UPDATE SET
		status = EXCLUDED.status,
		retry_count = EXCLUDED.retry_count,
		processed_at = EXCLUDED.processed_at;
	`

	_, err := r.db.Exec(
		query,
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		transaction.Currency,
		transaction.Status,
		transaction.RetryCount,
		transaction.MaxRetries,
		transaction.CreatedAt,
		transaction.ProcessedAt,
	)

	return err
}

func (r *PostgresTransactionRepository) FindByID(
	id uuid.UUID,
) (*domain.Transaction, error) {

	query := `
	SELECT
		id,
		account_id,
		amount,
		currency,
		status,
		retry_count,
		max_retries,
		created_at,
		processed_at
	FROM transactions
	WHERE id = $1
	`

	row := r.db.QueryRow(query, id)

	var transaction domain.Transaction

	err := row.Scan(
		&transaction.ID,
		&transaction.AccountID,
		&transaction.Amount,
		&transaction.Currency,
		&transaction.Status,
		&transaction.RetryCount,
		&transaction.MaxRetries,
		&transaction.CreatedAt,
		&transaction.ProcessedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *PostgresTransactionRepository) FindAll() ([]*domain.Transaction, error) {
	return nil, nil
}
