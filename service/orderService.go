package service

import (
	"StockMatchingEngine/models"
	"encoding/json"
	"fmt"
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

	matchingOrder(db)
}

//Get is responsible for retrieving the orders from the database
//and converting them to json objects
func (o *OrderService) Get(db *DatabaseService) ([]*models.Order, error) {
	rows, err := db.Query("SELECT id, userid, tickerid, price, quantity, command FROM orders ")
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

	for order := range orderBucket {
		json.Marshal(orderBucket[order])
	}

	return orderBucket, nil
}

//matchingOrder is responsible for matching the orders and passing them
//to the trade service so that the relevant fields in the database can
//be updated accordingly
func matchingOrder(db *DatabaseService) {
	orderService := OrderService{}
	orderBucket, err := orderService.Get(db)
	if err != nil {
		fmt.Println("An error occured while trying to get the order data")
	}

	BUY := []*models.Order{}
	SELL := []*models.Order{}

	for order := range orderBucket {
		if orderBucket[order].Command == "BUY  " {
			BUY = append(BUY, orderBucket[order])

		} else if orderBucket[order].Command == "SELL " {
			SELL = append(SELL, orderBucket[order])
		}
	}

	sort.Slice(BUY, func(i, j int) bool {
		return BUY[i].Price > BUY[j].Price
	})
	sort.Slice(SELL, func(i, j int) bool {
		return SELL[i].Price > SELL[j].Price
	})

	highestBuy := BUY[0]
	tradeMatches := [][]*models.Order{}
	var scenario string

	tradeService := &ServiceTrade{}
	for sellItem := range SELL {
		if highestBuy.Price >= SELL[sellItem].Price {
			if highestBuy.Quantity == SELL[sellItem].Quantity {
				scenario = "equal"
				tradeMatches = append(tradeMatches)
				fmt.Println("ENTER EQUAL ")
				trade := models.Trade{
					highestBuy.UserID,
					SELL[sellItem].UserID,
					SELL[sellItem].Price,
					SELL[sellItem].Quantity,
					SELL[sellItem].Ticker}

				tradeService.trade = &trade
				tradeService.Process(db, scenario, highestBuy.OrderID, SELL[sellItem].OrderID)

			} else if highestBuy.Quantity > SELL[sellItem].Quantity {
				scenario = "buyerMore"

				trade := models.Trade{
					highestBuy.UserID,
					SELL[sellItem].UserID,
					SELL[sellItem].Price,
					SELL[sellItem].Quantity,
					SELL[sellItem].Ticker}

				tradeService.trade = &trade
				tradeService.Process(db, scenario, highestBuy.OrderID, SELL[sellItem].OrderID)

			} else if highestBuy.Quantity < SELL[sellItem].Quantity {
				scenario = "sellerMore"

				trade := models.Trade{
					highestBuy.UserID,
					SELL[sellItem].UserID,
					SELL[sellItem].Price,
					highestBuy.Quantity,
					SELL[sellItem].Ticker}

				tradeService.trade = &trade
				tradeService.Process(db, scenario, highestBuy.OrderID, SELL[sellItem].OrderID)

			}
		}

	}
}
