package service

import (
	"StockMatchingEngine/model"
	"database/sql"
	"sort"
)

// OrderMatchingService provides access to the repository
type OrderMatchingService struct {
	OrderRepo model.Repository
}

// MatchingOrderEngine isresponsible for carrying out all of the actions
// from data retrieval to data preprocessing and update the relevant database
// fields by calling the update service
func (omg OrderMatchingService) MatchingOrderEngine(order *model.Order) error {
	queue, err := omg.getAndPrepareDataForMatching(order)
	if err != nil {
		return err
	}
	matchedOrderList, err := matchingOrder(queue, order)
	if err != nil {
		return err
	}

	updaterService := UpdaterService{omg.OrderRepo}
	updaterService.Update(matchedOrderList, order)

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

	orderQueueForMatch := populateQueues(orderBasket, order.Command)
	orderQueueForMatch = sortQueue(orderQueueForMatch)

	return orderQueueForMatch, nil

}

// MatchingOrder is responsible for matching the orders and  passing them
// to the trade service so that the relevant fields in the database can
// be updated accordingly
func matchingOrder(orderQueueForMatch []*model.Order, order *model.Order) ([]*model.Order, error) {
	var matchedOrders []*model.Order
	var (
		buy        = new(model.Order)
		sell       = new(model.Order)
		otherQueue model.Order
	)

	if order.Command == "BUY" {
		buy = order
		sell = &otherQueue
	} else if order.Command == "SELL" {
		buy = &otherQueue
		sell = order
	}

	for orderItem := range orderQueueForMatch {
		otherQueue = *orderQueueForMatch[orderItem]

		if buy.Price >= sell.Price {
			if buy.Quantity == sell.Quantity {
				matchedOrders = append(matchedOrders, sell)
				// matchedOrders = matchedOrders[1:] //remove first itm from queue

			} else if buy.Quantity > sell.Quantity {
				matchedOrders = append(matchedOrders, sell)

			} else if buy.Quantity < sell.Quantity {
				matchedOrders = append(matchedOrders, sell)

			}
		}
	}
	return matchedOrders, nil
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
func populateQueues(orderBucket []*model.Order, command string) []*model.Order {
	orderQueueForMatch := []*model.Order{}

	for order := range orderBucket {
		if orderBucket[order].Command != command {
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
