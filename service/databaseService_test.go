package service

import (
	"log"
	"os"
	"testing"
)

var a DatabaseService

func TestMain(m *testing.M) {
	a.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
	lastname TEXT NOT NULL,
    tickers INTEGER,
    trades INTEGER
)`
