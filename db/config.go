package db

import (
	"github.com/go-rel/rel/adapter/postgres"
	"log"
)

func init() {
	adapter, err := postgres.Open("postgres://postgres@localhost/rel_test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer adapter.Close()


}
