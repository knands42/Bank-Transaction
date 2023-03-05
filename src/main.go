package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/api"
	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/sqlc"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/util"
)

func main() {
	config := loadEnv()

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(*config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func loadEnv() *util.Config {
	config := util.NewConfig()
	path, _ := getRootFile()

	config.LoadConfig(path)

	return config
}

func getRootFile() (ex string, err error) {
	ex, _ = os.Getwd()
	_, err = os.Stat(filepath.Join(ex, "app.env"))

	if err != nil {
		ex = filepath.Join(ex, "../")
		_, err = os.Stat(filepath.Join(ex, "app.env"))

		if err != nil {
			log.Println("No env file provided, using only env variables")
		}
	}

	return
}
