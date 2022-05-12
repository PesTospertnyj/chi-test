package main

import (
	"chi-test/pkg/api/http"
	"chi-test/pkg/config"
	"chi-test/pkg/storage/postgres"
)

func main() {
	conf := config.New()

	db, err := postgres.New(conf)
	if err != nil {
		panic(err)
	}

	err = db.Seed()
	if err != nil {
		panic(err)
	}

	http.New(db, conf).Run()
}
