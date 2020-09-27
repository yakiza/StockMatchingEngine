package service

import (
	"StockMatchingEngine/models"
	"fmt"
)

type OrderService struct {
	*models.Order
}

//Create is responsible for inserting the Order into te database
func (o *OrderService) Create(db *DatabaseService) {
	s, err := db.Exec(
		"INSERT INTO orders(userid, tickerid, price, quantity, command) VALUES($1, $2, $3, $4, $5) ",
		o.UserID, o.Ticker, o.Price, o.Quantity, o.Command)
	fmt.Print(s)
	if err != nil {
		//Must change this
		fmt.Println("Something went wrong", err)
	}
}
