package handlers

import (
	"StockMatchingEngine/models"
	"StockMatchingEngine/service"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//	swagger:route GET /Users user userList
//	Returns a list of all the registered users saved in the database
//	responses:
//		200: User
func getUsers(rw http.ResponseWriter, r *http.Request) {

	db := service.DatabaseService{}
	db.InitializeDatabaseService(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	usrService := service.UserService{}
	userBaset, err := usrService.Get(&db) //Returns all users in JSON format

	if err != nil {
		fmt.Println("There was an error trying to get the data from the database\n", err)
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", userBaset)
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
