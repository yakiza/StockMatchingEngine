package storage_test

import (
	"StockMatchingEngine/model"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func DummyUser() model.User {
	return model.User{1, "Firstname", "Lastname"}

}

// TestCreateUser test the CreateUser method with a successful casse
func TestCreateUser(t *testing.T) {
	mock, repo := mockInit()

	user := DummyUser()

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Firstname, user.Lastname).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err := repo.CreateUser(&user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// // TestCreateUserFail tests the CreateUser with a failing case
func TestCreateUserFail(t *testing.T) {
	mock, repo := mockInit()

	user := DummyUser()

	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Firstname, user.Lastname).
		WillReturnError(fmt.Errorf("some error"))

	// now we execute our method
	if err := repo.CreateUser(&user); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
