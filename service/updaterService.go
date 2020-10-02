package service

import (
	"StockMatchingEngine/model"
)

type UpdaterService struct {
	OrderRepo model.Repository
}


//I am refactoring this NOW !
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
			us.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(trade)
			us.OrderRepo.CreateTickerOrSubstractQuantity(trade)
			us.OrderRepo.UpdateOrderQuantity(trade.BuyerID, trade.Quantity)
			us.OrderRepo.UpdateOrderQuantity(trade.SellerID, trade.Quantity)
			us.OrderRepo.CreateTrade(trade)
		} else if order.Quantity >= matchedOrdersList[orderMatched].Quantity {
			us.OrderRepo.CreateTrade(trade)
			if order.Command == "BUY" {
				us.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(trade)
				us.OrderRepo.CreateTickerOrSubstractQuantity(trade)
			} else if order.Command == "SELL" {

			}
		} else if order.Quantity <= matchedOrdersList[orderMatched].Quantity {
			if order.Command == "BUY" {

			} else if order.Command == "SELL" {

			}
			trade.Quantity = order.Quantity
			us.OrderRepo.CreateTrade(trade)
		}
	}

}
