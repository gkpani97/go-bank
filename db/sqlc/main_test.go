package db

import (
	"database/sql"
	"testing"
	"log"

	_  "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/go-bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = New(conn)
}