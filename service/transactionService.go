package service

import (
	"StockMatchingEngine/model"
	"fmt"
)

type BuyOrderbasket struct {
	Basket    []*model.Order
	OrderRepo model.Repository
}

// Transaction is responsible for carrying out the transactions for the matched
// orders that transactions includes:
// 1) Creation of trade
// 2) Substracting sold tickers
// 3) Adding bought tickers
// 4) Updating order quantities
func (b BuyOrderbasket) Transaction(currentOrder *model.Order) {

	for order := range b.Basket {
		trade := model.Trade{
			b.Basket[order].OrderID,
			currentOrder.OrderID,
			b.Basket[order].Price,
			b.Basket[order].Quantity,
			b.Basket[order].Ticker}
		b.OrderRepo.CreateTrade(&trade)
		b.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(&trade)
		b.OrderRepo.CreateTickerOrSubstractQuantity(&trade)
		b.OrderRepo.DecreaseOrderQuantity(b.Basket[order].OrderID, trade.Quantity)
		b.OrderRepo.DecreaseOrderQuantity(currentOrder.OrderID, trade.Quantity)

	}
}

type SellOrderbasket struct {
	Basket    []*model.Order
	OrderRepo model.Repository
}

// Transaction is responsible for carrying out the transactions for the matched
// orders that transactions includes:
// 1) Creation of trade
// 2) Substracting sold tickers
// 3) Adding bought tickers
// 4) Updating order quantities
func (s SellOrderbasket) Transaction(currentOrder *model.Order) {
	for order := range s.Basket {
		fmt.Println(s.Basket[order])
		fmt.Println("TRANSACTION SELL")

		trade := model.Trade{
			currentOrder.OrderID,
			s.Basket[order].OrderID,
			s.Basket[order].Price,
			s.Basket[order].Quantity,
			s.Basket[order].Ticker}
		s.OrderRepo.CreateTrade(&trade)
		s.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(&trade)
		s.OrderRepo.CreateTickerOrSubstractQuantity(&trade)
		s.OrderRepo.DecreaseOrderQuantity(s.Basket[order].OrderID, trade.Quantity)
		s.OrderRepo.DecreaseOrderQuantity(currentOrder.OrderID, trade.Quantity)
	}

}
