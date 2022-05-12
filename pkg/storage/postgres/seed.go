package postgres

import (
	"os"
)

func (p *postgres) Seed() error {
	file, err := os.ReadFile("cmd/seed/books.sql")
	if err != nil {
		return err
	}

	_, err = p.db.Exec(string(file))

	return err
}
