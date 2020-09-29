package service

import (
	"StockMatchingEngine/model"
	"log"
)

// ServiceTrade contains
type ServiceTrade struct {
	db *DatabaseService
}

// NewTradeService returns a new service with access to the database
func NewTradeService(db *DatabaseService) *ServiceTrade {
	return &ServiceTrade{db: db}
}

// Process takes updates the relevant fields on the database for the user and ticker table
// as wee as creates a new trade entry
func (st *ServiceTrade) Process(trade *model.Trade, scenario string, orderid ...int) error {

	log.Println("THe scenarios is -- >", scenario)
	log.Print("============================================")
	log.Print(len(scenario))
	log.Printf("%T\n", scenario)
	log.Printf(" %#v\n", []byte(scenario))
	log.Println(trade)

	_, err := st.db.Exec(
		"INSERT INTO trades(buyerid, sellerid, price, quantity, ticker) VALUES($1, $2, $3, $4, $5) ",
		trade.BuyerID, trade.SellerID, trade.Value, trade.Quantity, trade.Ticker)
	if err != nil {
		return err
	}

	log.Print("============================================ 1")

	// ADD INSERTION INTO THE TRADEBOOK PER	HAPS
	_, err = st.db.Exec(
		"INSERT INTO tickers (id, userid, quantity) VALUES($1, $2, $3) ON DUPLICATE KEY UPDATE id=$1, userid=$2, quantity= quantity + $3",
		trade.Ticker, trade.BuyerID, trade.Quantity)
	if err != nil {
		return err
	}
	log.Print("============================================ 2")

	_, err = st.db.Exec(
		"INSERT INTO tickers (id, userid, quantity) VALUES($1, $2, $3) ON DUPLICATE KEY UPDATE id=$1, userid=$2, quantity= quantity - $3",
		trade.Ticker, trade.SellerID, trade.Quantity)
	if err != nil {
		return err
	}
	log.Print("============================================ 3")

	// Updating accordingly the orders
	switch scenario {
	case "equal":
		log.Println("THe order ids that are being updated are==================>", orderid[0], orderid[1])
		_, err = st.db.Exec("UPDATE orders SET quantity=$1 WHERE id=$2 OR id=$3",
			0, orderid[0], orderid[1])
		if err != nil {
			return err
		}
	case "buyerMore":
		_, err = st.db.Exec("UPDATE orders SET quantity= quantity - $1 WHERE id=$2",
			trade.Quantity, orderid[0])
		if err != nil {
			return err
		}

		_, err = st.db.Exec("UPDATE orders SET quantity=$1 WHERE id=$2",
			0, orderid[1])
		if err != nil {
			return err
		}

	case "sellerMore":
		_, err = st.db.Exec("UPDATE orders SET quantity= 0 WHERE id=$2",
			trade.Quantity, orderid[0])
		if err != nil {
			return err
		}

		_, err = st.db.Exec("UPDATE orders SET quantity=quantity-$1 WHERE id=$2",
			0, orderid[1])

		if err != nil {
			return err
		}
	}

	return nil
}
