package storage

import "StockMatchingEngine/model"

//Create responsible for executing the sql query and creating a user
func (p PostgresOrderRepository) CreateUser(user *model.User) error {
	_, err := p.DB.Exec(
		`INSERT INTO
			users(firstname, lastname)
		VALUES
			($1, $2)`,
		user.Firstname, user.Lastname)

	return err
}

// // //Get  responsible for retriving the users from the database
// // func (p PostgresOrderRepository) GetAll() ([]*model.User, error) {
// // 	rows, err := u.db.Query("SELECT * FROM users")
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	defer rows.Close()

// // 	return users, nil
// // }
