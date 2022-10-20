package main

import (
	"database/sql"
	"log"

	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/api"
	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/sqlc"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
