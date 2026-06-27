CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,

    account_id VARCHAR(255) NOT NULL,

    amount NUMERIC(12,2) NOT NULL,

    currency VARCHAR(10) NOT NULL,

    status VARCHAR(50) NOT NULL,

    retry_count INTEGER NOT NULL,

    max_retries INTEGER NOT NULL,

    created_at TIMESTAMP NOT NULL,

    processed_at TIMESTAMP NULL
);