package service

import (
	"database/sql"
	"fmt"
	"log"
)

type DatabaseService struct {
	*sql.DB
}

//InitializeDatabaseService creates a connection to the database
func (d *DatabaseService) InitializeDatabaseService(user, password, dbname string) *sql.DB {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	d.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return d.DB
}
