package database

import (
	"database/sql"

	"github.com/ericgrandt/gqlgen-example/graph/model"
)

type UserData struct {
	db *sql.DB
}

func NewUserData(db *sql.DB) UserData {
	return UserData{
		db: db,
	}
}

func (data UserData) CreateUser(input model.NewUser) (model.User, error) {
	stmt, err := data.db.Prepare("INSERT INTO user(name) VALUES (?) RETURNING *")
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = stmt.QueryRow(input.Name).Scan(&user.ID, &user.Name)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (data UserData) GetUser(userID string) (model.User, error) {
	stmt, err := data.db.Prepare("SELECT * FROM user WHERE id = ?")
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = stmt.QueryRow(userID).Scan(&user.ID, &user.Name)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
