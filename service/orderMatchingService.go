package service

import (
	"StockMatchingEngine/model"
	"database/sql"
	"fmt"
	"sort"
)

// OrderMatchingService provides access to the repository
type OrderMatchingService struct {
	OrderRepo model.Repository
}

type BuyOrderMatchingType struct {
	order *model.Order
}

type SellOrderMatchingType struct {
	order *model.Order
}

// MatchingOrderEngine is responsible for carrying out all of the actions
// from data retrieval to data preprocessing and update the relevant database
// fields by calling the update service
func (omg OrderMatchingService) MatchingOrderEngine(order *model.Order) error {
	queue, err := omg.getAndPrepareDataForMatching(order)
	if err != nil {

		return err
	}

	if order.Command == "BUY" {
		buyOrderMatching := BuyOrderMatchingType{order}
		matchedOrderBasket, err := buyOrderMatching.matchingOrder(queue)
		if err != nil {
			return err
		}
		sellOrderBasket := SellOrderbasket{matchedOrderBasket, omg.OrderRepo}
		sellOrderBasket.Transaction(order)
	} else if order.Command == "SELL" {
		sellOrderMatching := SellOrderMatchingType{order}
		matchedOrderBasket, err := sellOrderMatching.matchingOrder(queue)
		if err != nil {
			return err
		}
		buyOrderBasket := BuyOrderbasket{matchedOrderBasket, omg.OrderRepo}
		buyOrderBasket.Transaction(order)
	}

	return nil
}

func (omg OrderMatchingService) getAndPrepareDataForMatching(order *model.Order) ([]*model.Order, error) {

	rows, err := omg.OrderRepo.GetActiveOrders(order.Ticker)
	if err != nil {
		return nil, err
	}

	// defer rows.Close()
	orderBasket, err := convertRowsToStruct(rows)
	if err != nil {
		return nil, err
	}

	orderQueueForMatch := populateQueues(orderBasket, order)
	orderQueueForMatch = sortQueue(orderQueueForMatch)

	return orderQueueForMatch, nil

}

// ConvertRowsToStruct converts the rows received in to stucts
func convertRowsToStruct(rows *sql.Rows) ([]*model.Order, error) {

	var orderBucket []*model.Order
	for rows.Next() {

		order := new(model.Order)
		err := rows.Scan(
			&order.OrderID,
			&order.UserID,
			&order.Ticker,
			&order.Price,
			&order.Quantity,
			&order.Command)
		if err != nil {
			return nil, err
		}
		orderBucket = append(orderBucket, order)
	}

	return orderBucket, nil

}

//Populates the corespinding queue
func populateQueues(orderBucket []*model.Order, currentOrder *model.Order) []*model.Order {
	orderQueueForMatch := []*model.Order{}

	for order := range orderBucket {
		// log.Printf("%T\n", currentOrder.Command)
		// log.Printf(" %#v\n", []byte(currentOrder.Command))
		// fmt.Println("--------------------------------------")
		// log.Printf("%T\n", orderBucket[order].Command)
		// log.Printf(" %#v\n", []byte(orderBucket[order].Command))
		if orderBucket[order].Command != currentOrder.Command {
			orderQueueForMatch = append(orderQueueForMatch, orderBucket[order])
		}
	}
	return orderQueueForMatch
}

// All algorithms in the Go sort package make O(n log n) comparisons
// in the worst case, where n is the number of elements to be sorted.
func sortQueue(queue []*model.Order) []*model.Order {

	sort.Slice(queue, func(i, j int) bool {
		return queue[i].Price > queue[j].Price
	})

	return queue
}

// MatchingOrder is responsible for matching the orders and  passing them
// to the trade service so that the relevant fields in the database can
// be updated accordingly
func (b BuyOrderMatchingType) matchingOrder(orderQueueForMatch []*model.Order) ([]*model.Order, error) {
	var matchedOrders []*model.Order
	for order := range orderQueueForMatch {
		if b.order.Price >= orderQueueForMatch[order].Price {
			if b.order.Quantity == orderQueueForMatch[order].Quantity {
				b.order.Quantity = 0
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
				fmt.Println(b.order.Price, "matched with", orderQueueForMatch[order])

			} else if b.order.Quantity > orderQueueForMatch[order].Quantity {
				b.order.Quantity -= orderQueueForMatch[order].Quantity
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
				fmt.Println(b.order, "matched with", orderQueueForMatch[order])
				orderQueueForMatch = orderQueueForMatch[1:] //remove first item from queue

			} else if b.order.Quantity < orderQueueForMatch[order].Quantity {
				orderQueueForMatch[order].Quantity -= b.order.Quantity
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
				fmt.Println(b.order, "matched with", orderQueueForMatch[order])
			}
		}
	}
	return matchedOrders, nil

}

// MatchingOrder is responsible for matching the orders and  passing them
// to the trade service so that the relevant fields in the database can
// be updated accordingly
func (s SellOrderMatchingType) matchingOrder(orderQueueForMatch []*model.Order) ([]*model.Order, error) {
	var matchedOrders []*model.Order
	for order := range orderQueueForMatch {
		if s.order.Quantity <= orderQueueForMatch[order].Quantity {
			if s.order.Quantity == orderQueueForMatch[order].Quantity {
				s.order.Quantity = 0
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
			} else if s.order.Quantity < orderQueueForMatch[order].Quantity {
				s.order.Quantity -= orderQueueForMatch[order].Quantity
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
				fmt.Println(s.order, "matched with", orderQueueForMatch[order])
			} else if s.order.Quantity > orderQueueForMatch[order].Quantity {
				orderQueueForMatch[order].Quantity -= s.order.Quantity
				matchedOrders = append(matchedOrders, orderQueueForMatch[order])
				fmt.Println(s.order, "matched with", orderQueueForMatch[order])
			}
		}
	}
	return matchedOrders, nil

}
