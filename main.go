package main

// context” refers to a package that provides functionality for request-scoped values and cancelation signals across API boundaries.

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/truont2/simplebank/api"
	db "github.com/truont2/simplebank/db/sqlc"
	"github.com/truont2/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Talk to Server:", err)
	}
}
