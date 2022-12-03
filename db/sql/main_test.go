package sql

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ndihey27/banking-app/utils"
	// "golang.org/x/tools/godoc/util"
)

var testQueries *Queries
var TestDbB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatal("Cannot Load Configuration: ", err)
	}

	TestDbB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	testQueries = New(TestDbB)

	os.Exit(m.Run())
}
