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

func (data UserData) CreateUser(user model.User) error {
	stmt, err := data.db.Prepare("INSERT INTO user(name) VALUES (?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name)
	if err != nil {
		return err
	}

	return nil
}
