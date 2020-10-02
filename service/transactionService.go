package service

import (
	"StockMatchingEngine/model"
)

type BUY struct {
	Basket []*model.Order
	OrderRepo model.Repository
}

type Trade struct {
	BuyerID  int     `json:"BuyerorderID"`
	SellerID int     `json:"sellerOrderID"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
	Ticker   string  `json:"ticker"`
}


func (b BUY) Transaction(currentOrder *model.Order) {
	for order := range b.Basket{
		trade := model.Trade{
				b.Basket[order].OrderID,
				currentOrder.OrderID,
				b.Basket[order].Price, 
				b.Basket[order].Quantity,
				b.Basket[order].Ticker}
		b.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(&trade)
		b.OrderRepo.CreateTickerOrSubstractQuantity(&trade)
		b.OrderRepo.UpdateOrderQuantity(trade.BuyerID, trade.Quantity)
		b.OrderRepo.UpdateOrderQuantity(trade.SellerID, trade.Quantity)
	}
}


type SELL struct{
	Basket []*model.Order
	OrderRepo model.Repository
}


func (s SELL) Transaction(currentOrder *model.Order) {
	for order := range s.Basket{
		trade := model.Trade{
				currentOrder.OrderID,
				s.Basket[order].OrderID,
				s.Basket[order].Price,
				s.Basket[order].Quantity,
				s.Basket[order].Ticker}
	s.OrderRepo.CreateTickerAddQuantityOrUpdateQuality(&trade)
	s.OrderRepo.CreateTickerOrSubstractQuantity(&trade)
	s.OrderRepo.UpdateOrderQuantity(trade.BuyerID, trade.Quantity)
	s.OrderRepo.UpdateOrderQuantity(trade.SellerID, trade.Quantity)

	}
}