package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gkpani97/go-bank/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	// config, err := util.LoadConfig("../..") //in ubuntu change the slash
	// config, err := util.LoadConfig("..\\..") // for windows
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	TestQueries = New(testDB)

	os.Exit(m.Run())
}
