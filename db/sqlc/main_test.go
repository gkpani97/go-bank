package db

import (
	"database/sql"
	"testing"
	"log"
	"os"

	_  "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/go-bank?sslmode=disable"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	TestQueries = New(testDB)

	os.Exit(m.Run())
}