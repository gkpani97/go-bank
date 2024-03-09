package db

import (
	"database/sql"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/gkpani97/go-bank/util"
	_ "github.com/lib/pq"
)

var TestQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var path string
	// config, err := util.LoadConfig("../..") //in ubuntu change the slash
	// config, err := util.LoadConfig("..\\..") // for windows

	os_type := runtime.GOOS

	if os_type == "windows" {
		path = "..\\.."
	} else {
		path = "../.."
	}

	config, err := util.LoadConfig(path)
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
