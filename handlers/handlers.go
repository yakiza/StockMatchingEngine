package handlers

import (
	"StockMatchingEngine/models"
	"StockMatchingEngine/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Handler is responsible to make calls to functions based on the url accessed
func Handler() http.Handler {

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", getUsers)
	getRouter.HandleFunc("/orders", getOrders)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/user/create", createUser)
	postRouter.HandleFunc("/order/create", createOrder)

	return router
}

//getUsers responsible for retrievig all the users saved in the database
func getUsers(rw http.ResponseWriter, r *http.Request) {

	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	usrService := service.UserService{}
	data, err := usrService.Get(&db)

	if err != nil {
		fmt.Println("There was an error trying to get the data from the database\n", err)
	}
	fmt.Fprintf(rw, "%s", data)
}

//createUser responsible for the creation
func createUser(rw http.ResponseWriter, r *http.Request) {
	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	usr := &models.User{}
	a := json.NewDecoder(r.Body)
	a.Decode(usr)
	usrService := service.UserService{usr}
	usrService.Create(&db)
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
	fmt.Println(order)

}
func createOrder(rw http.ResponseWriter, r *http.Request) {

}
