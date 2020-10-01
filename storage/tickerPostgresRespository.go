package storage

import (
	"StockMatchingEngine/model"
	"log"
)

//Here the function does 2 things but i think its okay
func (p PostgresOrderRepository) CreateTickerAddQuantityOrUpdateQuality(trade *model.Trade) error {
	// ADD INSERTION INTO THE TRADEBOOK PER	HAPS
	_, err := p.DB.SQL.Exec(
		`INSERT INTO
				tickers(id, userid, quantity)
			VALUES
				($1, $2, $3)
			ON DUPLICATE KEY UPDATE
				id=$1, userid=$2, quantity= quantity + $3`,
		trade.Ticker, trade.BuyerID, trade.Quantity)
	if err != nil {
		return err
	}
	return nil
}

// CreateTickerOrSubstractQuantity checks if there is an entry for the specific ticker and
// substracts OFCOURSE THERE SHOULD BE A CHECK IF THE USER HAS TICKERS TO SELLL BEFORE THE
// ORDER BUT THAT FALLS OUT OF THE SCOPE OF THIS APP SO THIS IS BEING USED INSTEAD.
func (p PostgresOrderRepository) CreateTickerOrSubstractQuantity(trade *model.Trade) error {

	_, err := p.DB.SQL.Exec(
		`INSERT INTO 
			tickers(id, userid, quantity) 
		VALUES
			($1, $2, $3) 
		ON DUPLICATE KEY UPDATE 
			id=$1, userid=$2, quantity= quantity - $3`,
		trade.Ticker, trade.SellerID, trade.Quantity)

	if err != nil {
		return err
	}
	return nil
}

// GetTickerLowerBuy queries the database for the min buy price and returns it
func (p PostgresOrderRepository) GetTickerLowerBuy(ticker string) (float64, error) {
	var lowestBuy float64

	err := p.DB.SQL.QueryRow("SELECT min(price) FROM orders WHERE quantity > 0 AND tickerid=$1 AND command='SELL'",
		ticker).Scan(&lowestBuy)
	if err != nil {
		log.Print("go")
		return 0, err
	}
	return lowestBuy, nil
}

// GetTickerHigherSell queries the database for the maximum sell price and returns it
func (p PostgresOrderRepository) GetTickerHigherSell(ticker string) (float64, error) {
	var highestSell float64

	err := p.DB.SQL.QueryRow("SELECT max(price) FROM orders WHERE quantity > 0 AND tickerid=$1 AND command='BUY'",
		ticker).Scan(&highestSell)
	if err != nil {
		return 0, err
	}
	return highestSell, nil
}
