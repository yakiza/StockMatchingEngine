package storage

import "StockMatchingEngine/model"

// CreateTrade is responsible to execute the query that adds a trade to the database
func (p PostgresOrderRepository) CreateTrade(trade *model.Trade) error {
	_, err := p.DB.Exec(
		`INSERT INTO
			trades(buyerid, sellerid, price, quantity, ticker)
		VALUES
			($1, $2, $3, $4, $5) `,
		trade.BuyerID, trade.SellerID, trade.Value, trade.Quantity, trade.Ticker)
	if err != nil {
		return err
	}
	return nil
}
