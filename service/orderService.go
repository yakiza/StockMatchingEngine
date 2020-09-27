package service

import (
	"StockMatchingEngine/models"
	"encoding/json"
	"fmt"
	"log"
	"sort"
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

//Get is responsible for retrieving the orders from the database
//and converting them to json objects
func (o *OrderService) Get(db *DatabaseService) ([]byte, error) {
	rows, err := db.Query("SELECT id, userid, tickerid, price, quantity, command FROM orders")
	if err != nil {
		fmt.Println("there was an error retrieving the data from the database", err)
	}
	defer rows.Close()
	orderBucket := make([]*models.Order, 0)

	fmt.Print(orderBucket)
	for rows.Next() {
		or := new(models.Order)

		err := rows.Scan(&or.OrderID, &or.UserID, &or.Ticker, &or.Price, &or.Quantity, &or.Command)
		if err != nil {
			fmt.Println("There was an error", err)
		}
		orderBucket = append(orderBucket, or)
	}
	b, err := json.Marshal(orderBucket)

	BUY := []*models.Order{}
	SELL := []*models.Order{}

	for i := range orderBucket {

		if orderBucket[i].Command == "BUY  " {
			BUY = append(BUY, orderBucket[i])

		} else if orderBucket[i].Command == "SELL " {
			SELL = append(SELL, orderBucket[i])

		}
		json.Marshal(orderBucket[i])

	}

	sort.Slice(BUY, func(i, j int) bool {
		return BUY[i].Price > BUY[j].Price
	})
	sort.Slice(SELL, func(i, j int) bool {
		return SELL[i].Price > SELL[j].Price
	})

	for i := range SELL {
		log.Print("Index", i, "Item", SELL[i], "PRICE ====>", SELL[i].Price)
	}
	for i := range BUY {
		log.Print("Index", i, "Item", BUY[i], "PRICE ====>", BUY[i].Price)
	}

	highestBuy := BUY[0]
	// tradeService := &ServiceTrade{}
	for sellItem := range SELL {
		if highestBuy.Price >= SELL[sellItem].Price {
			if highestBuy.Quantity >= SELL[sellItem].Quantity {

			}
		}
	}

	return b, nil
}

func matchingOrder() {
}
