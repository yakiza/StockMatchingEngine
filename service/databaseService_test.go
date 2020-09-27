package service

import (
	"StockMatchingEngine/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var a DatabaseService

func TestMain(m *testing.M) {
	a.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM users")
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

func TestEmptyTable(t *testing.T) {
	// clearTable()

	req, _ := http.NewRequest("GET", "/orders", nil)
	response := requestExecutor(req)

	responseCodeChecker(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func requestExecutor(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a := mux.NewRouter()
	router := handlers.Handler()
	router.ServeHTTP(rr, req)

	return rr
}

func responseCodeChecker(t *testing.T, expected, given int) {
	if expected != actual {
		t.Errorf("Expected response code %d. But instead got  %d\n", expected, given)
	}
}
