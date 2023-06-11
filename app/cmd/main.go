package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/api"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/db/migrations"
	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/db/sqlc"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/util"
)

func main() {
	config := loadEnv()
	dbConnection := connectDb(config)
	store := db.NewStore(dbConnection)
	hashingConfig := util.NewHashingConfig(config.PasswordHashSalt)

	serverInitializer(config, store, hashingConfig)
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

func serverInitializer(config *util.Config, store db.Store, hashingConfig *util.HashingConfig) {
	server, err := api.NewServer(*config, store, hashingConfig)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
