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
 
func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	TestQueries = New(conn)

	os.Exit(m.Run())
}