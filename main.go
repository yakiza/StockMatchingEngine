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
	// app := newApp()
	app := iris.New()

	db := new(service.DatabaseService)
	// service.NewDatabaseService(username, password, dbname)
	err := db.InitializeDatabaseService(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	if err != nil {
		app.Logger().Fatal(err)
	}

	app.PartyFunc("/", handlers.Router(db))
	addr := getAddr()
	app.Listen(addr)
}

// func newApp() *iris.Application {

// 	return app
// }

func getAddr() string {
	addr := ":8000"

	if v := os.Getenv("PORT"); v != "" {
		if v[0] != ':' {
			v = ":" + v
		}

		addr = v
	}

	return addr
}
