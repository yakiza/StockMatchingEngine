package model

//Order contains the attributes of the order data type
type Order struct {
	OrderID  int     `json:"orderID"`
	UserID   int     `json:"userID"`
	Ticker   string  `json:"ticker"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Command  string  `json:"command"`
}
