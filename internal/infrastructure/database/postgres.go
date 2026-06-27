package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/joaodddev/transaction-processing-engine/configs"
)

func NewPostgresConnection(
	cfg *configs.Config,
) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
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
