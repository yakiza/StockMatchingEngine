package storage

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"
	"database/sql"
)

//postgresOrder define the structure with a postgres transport type
type postgresOrder struct {
	OrderID  int     `sql:"orderid"`
	UserID   int     `sql:"userid"`
	Ticker   string  `sql:"ticker"`
	Price    float64 `sql:"price"`
	Quantity int     `sql:"quantity"`
	Command  string  `sql:"command"`
}

//PostgresOrderRepository provides access to the database
type PostgresOrderRepository struct {
	DB *service.DatabaseService
}

// NewPostgresOrderRepository Establishes a link
func NewPostgresOrderRepository(db *service.DatabaseService) *PostgresOrderRepository {
	if db == nil {
		panic("missing db")
	}
	return &PostgresOrderRepository{DB: db}
}

//CreateOrder performs the SQL query that inserts a new order into the database
func (p PostgresOrderRepository) CreateOrder(order *model.Order) error {
	_, err := p.DB.Exec(`
	INSERT INTO 
		orders (userid, tickerid, price, quantity, command) 
	VALUES ($1, $2, $3, $4, $5)`,
		order.UserID, order.Ticker, order.Price, order.Quantity, order.Command)

	if err != nil {
		return err
	}

	return nil
}

// GetActiveOrders is responsible for retrieving all the active orders and returning them
func (p PostgresOrderRepository) GetActiveOrders(ticker string) (*sql.Rows, error) {

	rows, err := p.DB.Query(
		`SELECT
			id, userid, tickerid, price, quantity, command
		FROM
			orders
		WHERE
			quantity > 0
		AND
			tickerid=$1 `, ticker)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// UpdateOrderQuantity is responsible for replacing the current quantity field to the
// quantity passed as an argument
func (p PostgresOrderRepository) UpdateOrderQuantity(orderID, quantity int) error {
	_, err := p.DB.Exec(`
		UPDATE
			orders
		SET
			quantity=quantity$1
		WHERE
			id=$2`,
		quantity, orderID)
	if err != nil {
		return err
	}
	return nil
}

// IncreaseOrderQuantity is responsible for increasing the current quantity by adding
// the quantity passed as an argument to the quantity stored already in the field
func (p PostgresOrderRepository) IncreaseOrderQuantity(orderID, quantity int) error {
	_, err := p.DB.Exec("UPDATE orders SET quantity= quantity + $1 WHERE id=$2",
		quantity, orderID)
	if err != nil {
		return err
	}
	return nil
}

// DecreaseOrderQuantity is responsible for decreasing the current quantity by substracting
// the quantity passed as an argument by the quantity stored in the field
func (p PostgresOrderRepository) DecreaseOrderQuantity(orderID, quantity int) error {

	_, err := p.DB.Exec("UPDATE orders SET quantity= quantity - $1 WHERE id=$2",
		quantity, orderID)
	if err != nil {
		return err
	}
	return nil
}
