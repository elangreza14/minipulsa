package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/elangreza14/minipulsa/authentication/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config := util.Config{
		DBDriver: "postgres",
		DBSource: "postgresql://minipulsa:minipulsa@localhost:5432/authentication?sslmode=disable",
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
