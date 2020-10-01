package model

import "database/sql"

// Repository defines interaction's with the storage stolution implemented,
// it provides access to the ORDER and USER storage.
type Repository interface {
	//Order Specific Methods
	CreateOrder(order *Order) error

	GetActiveOrders(ticker string) (*sql.Rows, error)

	UpdateOrderQuantity(orderID, quantity int) error

	IncreaseOrderQuantity(orderID, quantity int) error

	DecreaseOrderQuantity(orderID, quantity int) error

	CreateUser(user *User) error

	//Ticker Specific Methods
	CreateTickerOrSubstractQuantity(trade *Trade) error

	CreateTickerAddQuantityOrUpdateQuality(trade *Trade) error

	GetTickerLowerBuy(ticker string) (float64, error)

	GetTickerHigherSell(ticker string) (float64, error)

	//Trade Specific Methods
	CreateTrade(trade *Trade) error
}
