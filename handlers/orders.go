package handlers

import (
	"StockMatchingEngine/models"
	"StockMatchingEngine/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func createOrder(rw http.ResponseWriter, r *http.Request) {

	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	order := &models.Order{}
	a := json.NewDecoder(r.Body)
	a.Decode(order)
	orderService := service.OrderService{order}
	orderService.Create(&db)
	fmt.Println(order)

}

func getOrders(rw http.ResponseWriter, r *http.Request) {
	//Must find how to overcome this (pass as argument)
	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	order := &models.Order{}
	a := json.NewDecoder(r.Body)
	a.Decode(order)
	orderService := service.OrderService{order}
	orderService.Create(&db)
	// b, err := json.Marshal(orderBucket)

}
