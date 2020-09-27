package service

import (
	"StockMatchingEngine/models"
	"encoding/json"
	"fmt"
)

type UserService struct {
	*models.User
}

//Get  responsible for retriving the users from the database
func (u *UserService) Get(db *DatabaseService) ([]byte, error) {

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("there was an error retrieving the data from the database")
	}
	defer rows.Close()
	usrs := make([]*models.User, 0)
	for rows.Next() {
		ur := new(models.User)
		err := rows.Scan(&ur.ID, &ur.Firstname, &ur.Lastname, &ur.Ticker, &ur.Trades)
		if err != nil {
			fmt.Println("There was an error", err)
		}
		usrs = append(usrs, ur)
	}

	b, err := json.Marshal(usrs[0])

	return b, nil
}

//Create responsible for executing the sql query and creating a user
func (u *UserService) Create(db *DatabaseService) error {
	s, err := db.Exec(
		"INSERT INTO users(firstname, lastname) VALUES($1, $2) ",
		u.Firstname, u.Lastname)
	fmt.Print(s)
	if err != nil {
		//Must change this
		fmt.Println("IT did not insert the shit in to the db", err)
		return nil
	}

	return nil

}
