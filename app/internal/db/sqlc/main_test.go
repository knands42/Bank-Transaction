package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/db/migrations"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config := util.NewConfig()
	err := config.LoadConfig("test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	migrateDir, err := util.GetMigrationsFolder()
	if err != nil {
		log.Fatal("cannot get migrations folder:", err)
	}
	migrations.Up(testDB, os.DirFS(migrateDir))
	testQueries = New(testDB)

	defer migrations.Down(testDB, os.DirFS(migrateDir))
	os.Exit(m.Run())
}
