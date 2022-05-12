package postgres

import (
	"database/sql"

	"chi-test/pkg/config"
)

type postgres struct {
	db *sql.DB
}

func New(config *config.Config) (*postgres, error) {
	db, err := connect(config)
	if err != nil {
		return nil, err
	}

	return &postgres{db: db}, nil
}
