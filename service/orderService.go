package service

import (
	"fmt"
	"log"
	"sort"

	"StockMatchingEngine/model"
)

type OrderService struct {
	db *DatabaseService
}

func NewOrderService(db *DatabaseService) *OrderService {
	return &OrderService{
		db: db,
	}
}

// Create is responsible for inserting the Order into te database
func (o *OrderService) Create(order *model.Order, tradeService *ServiceTrade) error {
	log.Println("=================================== - (1A)")

	var currentOrderID int
	var currentOrderTicker string
	var currentORderCommand string
	err := o.db.QueryRow("INSERT INTO orders(userid, tickerid, price, quantity, command) VALUES($1, $2, $3, $4, $5) RETURNING id, tickerid, command ",
		order.UserID, order.Ticker, order.Price, order.Quantity, order.Command).Scan(&currentOrderID, &currentOrderTicker, &currentORderCommand)
	if err != nil {
		return err
	}

	MatchingOrder(o, tradeService, currentOrderID, currentOrderTicker) ///passing the id of the object that was just created and the ticker symbol
	return nil
}

// GetAll returns all available orders.
func (o *OrderService) GetAll() ([]*model.Order, error) {
	rows, err := o.db.Query("SELECT id, userid, tickerid, price, quantity, command FROM orders WHERE quantity > 0 ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderBucket []*model.Order

	for rows.Next() {
		or := new(model.Order)

		err := rows.Scan(&or.OrderID, &or.UserID, &or.Ticker, &or.Price, &or.Quantity, &or.Command)
		if err != nil {
			return nil, err
		}
		orderBucket = append(orderBucket, or)
	}

	return orderBucket, nil
}

// GetAllActiveOrders is responsible for retrieving the orders from the database
// and converting them to json objects
func (o *OrderService) GetAllActiveOrders(ticker string) ([]*model.Order, error) {
	rows, err := o.db.Query("SELECT id, userid, tickerid, price, quantity, command FROM orders WHERE quantity > 0 AND tickerid=$1 ", ticker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderBucket []*model.Order

	for rows.Next() {
		or := new(model.Order)

		err := rows.Scan(&or.OrderID, &or.UserID, &or.Ticker, &or.Price, &or.Quantity, &or.Command)
		if err != nil {
			return nil, err
		}
		orderBucket = append(orderBucket, or)
	}

	return orderBucket, nil
}

// GetLowerBuyHigherSell queries the database for the min and max values
// stored in the price field and returns them
func (o *OrderService) GetLowerBuyHigherSell(ticker string) (float64, float64, error) {
	var (
		lowestBuy   float64
		highestSell float64
	)

	err := o.db.QueryRow("SELECT min(price) FROM orders WHERE quantity > 0 AND tickerid=$1 AND command='SELL'",
		ticker).Scan(&lowestBuy)
	if err != nil {
		log.Print("go")
		return 0, 0, err
	}

	err = o.db.QueryRow("SELECT max(price) FROM orders WHERE quantity > 0 AND tickerid=$1 AND command='BUY'",
		ticker).Scan(&highestSell)
	if err != nil {
		return 0, 0, err
	}

	return lowestBuy, highestSell, nil
}

// MatchingOrder is responsible for matching the orders and  passing them
// to the trade service so that the relevant fields in the database can
// be updated accordingly
func MatchingOrder(orderService *OrderService, tradeService *ServiceTrade, orderId int, ticker string) error {

	orderBucket, err := orderService.GetAllActiveOrders(ticker) // retrieving a list of BUY & SELL orers to populate the queues
	if err != nil {
		return err
	}

	//Populating and sorting Queues
	unSortedBUY, unSortedSELL := populateQueues(orderBucket)
	buyQueue, sellQueue := sortQueues(unSortedBUY, unSortedSELL)

	highestBuy := buyQueue[0]
	var scenario string
	var tradeVal model.Trade

	var trade *model.Trade = &tradeVal

	for sellItem := range sellQueue {
		if highestBuy.Price >= sellQueue[sellItem].Price {
			log.Print(highestBuy.Price, " is equal to ", sellQueue[sellItem].Price)
			*trade = model.Trade{
				highestBuy.UserID,
				sellQueue[sellItem].UserID,
				sellQueue[sellItem].Price,
				sellQueue[sellItem].Quantity,
				sellQueue[sellItem].Ticker}

			fmt.Println("TRADE object before", *trade)
			if highestBuy.Quantity == sellQueue[sellItem].Quantity {
				scenario = "equal"

				tradeService.Process(trade, scenario, highestBuy.OrderID, sellQueue[sellItem].OrderID)
				fmt.Println("Matched ", highestBuy, sellQueue[sellItem])

			} else if highestBuy.Quantity > sellQueue[sellItem].Quantity {
				scenario = "buyerMore"

				tradeService.Process(trade, scenario, highestBuy.OrderID, sellQueue[sellItem].OrderID)
				fmt.Println("Matched ", highestBuy, sellQueue[sellItem])

			} else if highestBuy.Quantity < sellQueue[sellItem].Quantity {
				scenario = "sellerMore"

				trade.Quantity = highestBuy.Quantity
				tradeService.Process(trade, scenario, highestBuy.OrderID, sellQueue[sellItem].OrderID)
				fmt.Println("Matched ", highestBuy, sellQueue[sellItem])
			}
		}
	}

	return nil
}

func populateQueues(orderBucket []*model.Order) ([]*model.Order, []*model.Order) {
	BUY := []*model.Order{}
	SELL := []*model.Order{}

	for order := range orderBucket {
		if orderBucket[order].Command == "BUY  " {
			BUY = append(BUY, orderBucket[order])

		} else if orderBucket[order].Command == "SELL " {
			SELL = append(SELL, orderBucket[order])
		}
	}
	return BUY, SELL
}

func sortQueues(queue ...[]*model.Order) ([]*model.Order, []*model.Order) {

	BUY := queue[0]
	SELL := queue[1]
	//Sorting queyes so that BUY[0] contains the highest offer and the SELL[0] contains the less value
	sort.Slice(BUY, func(i, j int) bool {
		return BUY[i].Price > BUY[j].Price
	})
	sort.Slice(SELL, func(i, j int) bool {
		return SELL[i].Price > SELL[j].Price
	})

	return BUY, SELL
}
