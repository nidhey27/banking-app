package sql

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:mypassword@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var TestDbB *sql.DB

func TestMain(m *testing.M) {
	var err error
	TestDbB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	testQueries = New(TestDbB)

	os.Exit(m.Run())
}
