package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=localhost port=5433 user=postgres password=postgres dbname=transaction_engine sslmode=disable",
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
