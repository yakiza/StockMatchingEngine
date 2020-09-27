package service

import (
	"StockMatchingEngine/models"
	"fmt"
	"log"
)

type ServiceTrade struct {
	trade *models.Trade
}

func (st *ServiceTrade) Process(db *DatabaseService, scenario string, orderid ...int) {
	fmt.Println("Beggining of process")
	_, err := db.Exec(
		"INSERT INTO trades(buyerid, sellerid, price, quantity, ticker) VALUES($1, $2, $3, $4, $5) ",
		st.trade.BuyerID, st.trade.SellerID, st.trade.Value, st.trade.Quantity, st.trade.Ticker)
	// ADD INSERTION INTO THE TRADEBOOK PER	HAPS
	db.Exec(
		"INSERT INTO tickers (id, userid, quantity) VALUES($1, $2, $3) ON DUPLICATE KEY UPDATE id=$1, userid=$2, quantity= quantity + $3",
		st.trade.Ticker, st.trade.BuyerID, st.trade.Quantity)
	db.Exec(
		"INSERT INTO tickers (id, userid, quantity) VALUES($1, $2, $3) ON DUPLICATE KEY UPDATE id=$1, userid=$2, quantity= quantity - $3",
		st.trade.Ticker, st.trade.SellerID, st.trade.Quantity)

	if err != nil {
		fmt.Println("----> Something went wrong \n", err)
	}

	// Updating accordingly the orders
	switch scenario {

	case "equal":
		fmt.Println("Beggining of process")

		db.Exec("UPDATE orders SET quantity=$1 WHERE id=$2 OR id=$3",
			0, orderid[0], orderid[1])

	case "buyerMore":
		db.Exec("UPDATE orders SET quantity= quantity + $1 WHERE id=$2",
			st.trade.Quantity, orderid[0])

		db.Exec("UPDATE orders SET quantity=$1 WHERE id=$2",
			st.trade.Quantity, orderid[1])

	case "sellerMore":
		db.Exec("UPDATE orders SET quantity= quantity - $1 WHERE id=$2",
			st.trade.Quantity, orderid)

		db.Exec("UPDATE orders SET quantity=$1 WHERE id=$2",
			0, orderid)

	}
	log.Print("TRADE COMPLETED")

}
