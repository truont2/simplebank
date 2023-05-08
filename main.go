package main

// context” refers to a package that provides functionality for request-scoped values and cancelation signals across API boundaries.

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/truont2/simplebank/api"
	db "github.com/truont2/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot Talk to Server:", err)
	}
}
