package storage_test

import (
	"StockMatchingEngine/model"
	"StockMatchingEngine/service"
	"StockMatchingEngine/storage"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// Exec mock of the exec function of pg library
// func (mock *MockDatabaseService) Exec(query string) (interface{}, error) {
// 	args := mock.Called(query)
// 	return nil, args.Error(1)
// }

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreateOrder(t *testing.T) {
	db, mock := NewMock()
	serv := &service.DatabaseService{db}
	repo := storage.PostgresOrderRepository{serv}
	defer func() {
		repo.DB.SQL.Close()
	}()

	order := model.Order{1, 1, "AABB", 50.50, 50, "SELL"}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO orders").
		WithArgs(order.UserID, order.Ticker, order.Price, order.Quantity, order.Command)
	mock.ExpectCommit()

	// now we execute our method
	if err := repo.CreateOrder(&order); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// func TestCreateOrderFail(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &storage.PostgresOrderRepository{db}
// 	defer func() {
// 		repo.DB.Close()
// 	}()
// 	order := model.Order{}

// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO orders").
// 		WithArgs().
// 		WillReturnError(fmt.Errorf("some error"))
// 	mock.ExpectRollback()

// 	// now we execute our method
// 	if err := repo.CreateOrder(&order); err == nil {
// 		t.Errorf("was expecting an error, but there was none")
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

// }
// func TestCreateOrderFail(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &storage.PostgresOrderRepository{db}
// 	defer func() {
// 		repo.DB.Close()
// 	}()
// 	order := model.Order{}
// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO orders").
// 		WithArgs(5).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// }

// 	db := new(MockDatabaseService)
// 	OrderRepo := storage.NewPostgresOrderRepository(db) //arrange  -> preparation
// 	order := model.Order{1, 1, "AABB", 50.50, 50, "SELL"}
// 	db.On("Exec").Return(nil, nil)

// 	err := OrderRepo.CreateOrder(order) // act

// 	assert.Nil(err) //
// }

// func TestCreateOrderFail(t *testing.T) {
// 	db := new(MockDatabaseService)
// 	OrderRepo := storage.NewPostgresOrderRepository(db) //arrange  -> preparation
// 	order := model.Order{1, 1, "AABB", 50.50, 50, "SELL"}
// 	db.On("Exec").Return(nil, error)

// 	err := OrderRepo.CreateOrder(order) // act

// assert.Error(err)
// }
