package storage_test

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"
	"StockMatchingEngine/storage"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
func mockInit() (sqlmock.Sqlmock, storage.PostgresOrderRepository) {
	db, mock := NewMock()
	serv := &service.DatabaseService{db}
	repo := storage.PostgresOrderRepository{serv}

	return mock, repo
}

func TestCreateOrder(t *testing.T) {
	mock, repo := mockInit()

	order := model.Order{1, 1, "AABB", 50.50, 50, "SELL"}
	mock.ExpectExec("INSERT INTO orders").
		WithArgs(order.UserID, order.Ticker, order.Price, order.Quantity, order.Command).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err := repo.CreateOrder(&order); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateOrderFail(t *testing.T) {
	mock, repo := mockInit()

	order := model.Order{1, 1, "AABB", 50.50, 50, "SELL"}
	mock.ExpectExec("INSERT INTO orders").
		WithArgs(order.UserID, order.Ticker, order.Price, order.Quantity, order.Command).
		WillReturnError(fmt.Errorf("some error"))

	// now we execute our method
	if err := repo.CreateOrder(&order); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// }
// func TestGetActiveOrders(t *testing.T) {
// 	mock, repo := mockInit()
// 	ticker := "AABB"
// 	mock.ExpectExec(`SELECT
// 		id, userid, tickerid, price, quantity, command
// 		FROM
// 			orders
// 		WHERE
// 			quantity > 0
// 		AND
// 			tickerid=$1 `).
// 		WithArgs(ticker)

// 	// now we execute our method
// 	if _, err := repo.GetActiveOrders(ticker); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// }

// func TestGetActiveOrdersFail(t *testing.T) {
// 	mock, repo := mockInit()
// 	// now we execute our method
// 	if err := repo.CreateOrder(&order); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestUpdateOrderQuantity(t *testing.T) {
// 	mock, repo := mockInit()

// 	// now we execute our method
// 	if err := repo.CreateOrder(&order); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// }

// func TestUpdateOrderQuantityFail(t *testing.T) {
// 	mock, repo := mockInit()
// 	// now we execute our method
// 	if err := repo.CreateOrder(); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestIncreaseOrderQuantity(t *testing.T) {
// 	mock, repo := mockInit()

// 	// now we execute our method
// 	if err := repo.CreateOrder(&order); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
// func TestIncreaseOrderQuantityFail(t *testing.T) {
// 	mock, repo := mockInit()
// 	// now we execute our method
// 	if err := repo.CreateOrder(); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestDecreaseOrderQuantity(t *testing.T) {
// 	mock, repo := mockInit()

// 	// now we execute our method
// 	if err := repo.CreateOrder(&order); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestDecreaseOrderQuantityFail(t *testing.T) {
// 	mock, repo := mockInit()

// 	// now we execute our method
// 	if err := repo.CreateOrder(); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
