package db_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/src/infrastructure/db/sqlc"
	_ "github.com/lib/pq"
)

var testQueries *db.Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(conn)

	os.Exit(m.Run())
}
