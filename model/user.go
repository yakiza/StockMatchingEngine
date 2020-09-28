package model

//User contains the attributes of the user data type
type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Ticker    string `json:"ticker"`
	Trades    string `json:"trades"`
}
