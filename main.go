//	Stock Matching Engine
//
//	Documentation for Stock Matching API
//
//	Scheme: http
//	BasePath: /order/
//  Version: 0.0.1
// swagger:meta
package main

import (
	"os"

	"StockMatchingEngine/handlers"
	"StockMatchingEngine/service"

	"github.com/kataras/iris/v12"
)

func main() {
	app := newApp()

	addr := getAddr()
	app.Listen(addr)
}

func newApp() *iris.Application {

	db := new(service.DatabaseService)
	app := iris.New()

	// service.NewDatabaseService(username, password, dbname)
	err := db.InitializeDatabaseService(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.PartyFunc("/", handlers.Router(db))

	if err != nil {
		app.Logger().Fatal(err)
	}

	return app
}

func getAddr() string {
	addr := ":8080"

	if v := os.Getenv("PORT"); v != "" {
		if v[0] != ':' {
			v = ":" + v
		}

		addr = v
	}

	return addr
}
