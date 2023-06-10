package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/api"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/migrations"
	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/sqlc"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/util"
)

func main() {
	config := loadEnv()
	dbConnection := connectDb(config)
	store := db.NewStore(dbConnection)

	serverInitializer(config, store)
}

func loadEnv() *util.Config {
	config := util.NewConfig()
	config.LoadConfig(config.Profile)
	return config
}

func connectDb(config *util.Config) *sql.DB {
	var err error
	dbConnection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	migrateDir, err := util.GetMigrationsFolder()
	if err != nil {
		log.Fatal("cannot get migrations folder:", err)
	}
	migrations.Up(dbConnection, os.DirFS(migrateDir))

	return dbConnection
}

func serverInitializer(config *util.Config, store db.Store) {
	server, err := api.NewServer(*config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
