package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DatabaseService dependency
type DatabaseService struct {
	*sql.DB
}

//InitializeDatabaseService creates a connection to the database
func (d *DatabaseService) InitializeDatabaseService(host, port, user, password, dbname string) error {
	connectionString :=
		fmt.Sprintf(" user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
		// host=%s port=%shost, port,
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
