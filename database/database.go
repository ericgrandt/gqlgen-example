package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetOrCreateDatabase(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", dbName))
	if err != nil {
		log.Fatal(err)
	}

	createUserTable(db)
	createTodoTable(db)
	createTagTable(db)
	createTodoTagTable(db)

	return db
}

func createUserTable(db *sql.DB) {
	stmt := `
    CREATE TABLE IF NOT EXISTS user (
        id INT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL
    );
    `

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func createTodoTable(db *sql.DB) {
	stmt := `
    CREATE TABLE IF NOT EXISTS todo (
        id INT NOT NULL PRIMARY KEY,
		user_id int NOT NULL,
		value VARCHAR NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id)
    );
    `

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func createTagTable(db *sql.DB) {
}

func createTodoTagTable(db *sql.DB) {
}
