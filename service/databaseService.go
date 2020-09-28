package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseService struct {
	*sql.DB
}

//InitializeDatabaseService creates a connection to the database
func (d *DatabaseService) InitializeDatabaseService(host, port, user, password, dbname string) error {
	// host=%s port=%shost, port,
	connectionString :=
		fmt.Sprintf(" user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	d.DB = db
	return nil
}
