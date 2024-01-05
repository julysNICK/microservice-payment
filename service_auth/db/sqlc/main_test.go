package db

import (
	"database/sql"
	"os"
	"service_auth/util/config"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../")

	if err != nil {
		panic(err)
	}

	testDB, err = sql.Open("postgres", config.DSN)

	if err != nil {
		panic(err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
