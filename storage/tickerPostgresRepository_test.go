package storage_test

import (
	"StockMatchingEngine/model"
)

func dummyTicker() model.Ticker {
	return model.Ticker{"AABB"}
}

// func TestCreateTickerAddQuantityOrUpdateQuality(t *testing.T) {
// 	mock, repo := mockInit()

// 	trade := dummyTrade()
// 	mock.ExpectExec("INSERT INTO tickers VALUES (.+) ON DUPLICATE KEY UPDATE (.+)").
// 		WithArgs(trade.Ticker, trade.BuyerID, trade.Quantity).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// now we execute our method
// 	if err := repo.CreateTickerAddQuantityOrUpdateQuality(&trade); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestCreateTickerAddQuantityOrUpdateQualityFail(t *testing.T) {
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

// func TestCreateTickerOrSubstractQuantity(t *testing.T) {
// 	mock, repo := mockInit()

// 	// now we execute our method
// 	if err := repo.CreateOrder(); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }
// func TestCreateTickerOrSubstractQuantityFail(t *testing.T) {
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

// func TestGetTickerLowerBuy(t *testing.T) {
// 	mock, repo := mockInit()

// 	ticker := "AABB"
// 	mock.ExpectExec("SELECT min((.+)) FROM orders Where quantity > 0 AND tickerid=(.+) AND command='SELL'").
// 		WithArgs(ticker).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// now we execute our method
// 	if _, err := repo.GetTickerLowerBuy(ticker); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}
// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// func TestGetTickerLowerBuyFail(t *testing.T) {
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
// func TestGetTickerHigherSell(t *testing.T) {
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

// func TestGetTickerHigherSellFail(t *testing.T) {
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
