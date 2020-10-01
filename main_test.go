package main

// import (
// 	"encoding/json"
// 	"log"
// 	"os"
// 	"testing"

// 	"StockMatchingEngine/model"
// 	"StockMatchingEngine/service"

// 	"github.com/kataras/iris/v12/httptest"
// )

// var a service.DatabaseService

// func TestMain(m *testing.M) {
// 	a.InitializeDatabaseService(
// 		os.Getenv("APP_DB_HOST"),
// 		os.Getenv("APP_DB_PORT"),
// 		os.Getenv("APP_DB_USERNAME"),
// 		os.Getenv("APP_DB_PASSWORD"),
// 		os.Getenv("APP_DB_NAME"))

// 	ensureTableExists()
// 	code := m.Run()
// 	os.Exit(code)
// }

// //ensureTableExists tries to execute the sql query in order to check if database connection exists
// func ensureTableExists() {
// 	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}
// }

// const tableCreationQuery = `
// CREATE TABLE IF NOT EXISTS users
// (
//     id SERIAL PRIMARY KEY,
//     firstname TEXT NOT NULL,
// 	lastname TEXT NOT NULL
// )`

// //Testing the API  functionality of creation and retireving data
// func TestStockMatchingEngineAPI(t *testing.T) {
// 	app := newApp()

// 	e := httptest.New(t, app)

// 	// var user = map[string]interface{}{
// 	// 	"firstname": "john",
// 	// 	"lastname":  "doe",
// 	// }

// 	// var order = map[string]interface{}{
// 	// 	"userID":   "1",
// 	// 	"ticker":   "AABB",
// 	// 	"price":    "1",
// 	// 	"quantity": "2",
// 	// 	"command":  "BUY",
// 	// }

// 	user := model.User{1, "john", "doe"}
// 	userJSON, err := json.Marshal(user)
// 	if err != nil {
// 		return
// 	}

// 	orderSELL := model.Order{1, 1, "AABB", 1.1, 5, "SELL"}
// 	orderSELLJSON, err := json.Marshal(orderSELL)
// 	if err != nil {
// 		return
// 	}

// 	orderBUY := model.Order{1, 1, "AABB", 1.1, 5, "SELL"}
// 	orderBUYJSON, err := json.Marshal(orderBUY)
// 	if err != nil {
// 		return
// 	}

// 	// test create user.
// 	e.POST("/users").WithJSON(userJSON).Expect().Status(httptest.StatusCreated)
// 	// test get the created user.
// 	e.GET("/users").Expect().Status(httptest.StatusOK).JSON().Array().Contains(
// 		user)

// 	// test create orders.
// 	e.POST("/orders").WithJSON(orderSELLJSON).Expect().Status(httptest.StatusCreated)
// 	e.POST("/orders").WithJSON(orderBUYJSON).Expect().Status(httptest.StatusCreated)

// 	// test get the created orders.
// 	e.GET("/orders").Expect().Status(httptest.StatusOK).JSON().Array().Contains(
// 		orderSELLJSON, orderBUYJSON)
// }
