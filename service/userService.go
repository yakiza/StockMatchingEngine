package service

import (
	"StockMatchingEngine/model"
)

type UserService struct {
	db *DatabaseService
}

func NewUserService(db *DatabaseService) *UserService {
	return &UserService{
		db: db,
	}
}

//Get  responsible for retriving the users from the database
func (u *UserService) GetAll() ([]*model.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

//Create responsible for executing the sql query and creating a user
func (u *UserService) Create(user *model.User) error {
	_, err := u.db.Exec(
		`INSERT INTO users (firstname, lastname) VALUES ($1, $2)`,
		user.Firstname, user.Lastname)
	return err
}
