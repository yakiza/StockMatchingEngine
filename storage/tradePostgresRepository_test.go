package storage_test

import (
	"StockMatchingEngine/model"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func dummyTrade() model.Trade {
	return model.Trade{1, 1, 2.2, 1, "AABB"}
}

func TestCreateTrade(t *testing.T) {
	mock, repo := mockInit()

	trade := dummyTrade()

	mock.ExpectExec("INSERT INTO trades").
		WithArgs(trade.BuyerID, trade.SellerID, trade.Value, trade.Quantity, trade.Ticker).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err := repo.CreateTrade(&trade); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTradeFail(t *testing.T) {
	mock, repo := mockInit()

	trade := dummyTrade()
	mock.ExpectExec("INSERT INTO trades").
		WithArgs(trade.BuyerID, trade.SellerID, trade.Value, trade.Quantity, trade.Ticker).
		WillReturnError(fmt.Errorf("some error"))

	// now we execute our method
	if err := repo.CreateTrade(&trade); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
