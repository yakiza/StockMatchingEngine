package models

//Trade structure describing the trade attributes
type Trade struct {
	BuyerID  int     `json:"BuyerorderID"`
	SellerID int     `json:"sellerOrderID"`
	Value    float64 `json:"value"`
	Quantity int     `json:"quantity"`
	Ticker   string  `json:"ticker"`
}
