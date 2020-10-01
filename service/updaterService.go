package service

import (
	"StockMatchingEngine/model"
)

type UpdaterService struct {
	OrderRepo model.Repository
}

func (us UpdaterService) Update(matchedOrdersList []*model.Order, order *model.Order) {
	var tradeVal model.Trade
	var trade *model.Trade = &tradeVal

	for orderMatched := range matchedOrdersList {
		*trade = model.Trade{
			order.OrderID,
			matchedOrdersList[orderMatched].UserID,
			matchedOrdersList[orderMatched].Price,
			matchedOrdersList[orderMatched].Quantity,
			matchedOrdersList[orderMatched].Ticker}
		if order.Quantity == matchedOrdersList[orderMatched].Quantity {
			us.OrderRepo.CreateTrade(trade)
			us.OrderRepo.UpdateOrderQuantity(order.OrderID, order.Quantity)
			us.OrderRepo.UpdateOrderQuantity(matchedOrdersList[orderMatched].OrderID, order.Quantity)
		} else if order.Quantity >= matchedOrdersList[orderMatched].Quantity {
			us.OrderRepo.CreateTrade(trade)
			// us.OrderRepo.IncreaseOrderQuantity()
			// us.OrderRepo.DecreaseOrderQuantity()
		} else if order.Quantity <= matchedOrdersList[orderMatched].Quantity {
			trade.Quantity = order.Quantity
			us.OrderRepo.CreateTrade(trade)
			// us.OrderRepo.IncreaseOrderQuantity()
			// us.OrderRepo.DecreaseOrderQuantity()
		}
	}

}
