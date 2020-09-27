package main

import (
	"StockMatchingEngine/handlers"
	service "StockMatchingEngine/service"
	"log"
	"net/http"
	"os"
)

func main() {
	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	router := handlers.Handler()

	log.Print("Listening at address 127.0.0.1:8000")
	http.ListenAndServe(":8000", router)
}
