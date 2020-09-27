package handlers

import (
	"net/http"

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
