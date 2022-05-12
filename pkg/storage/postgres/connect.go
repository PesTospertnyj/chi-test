package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"chi-test/pkg/config"
)

func connect(config *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
