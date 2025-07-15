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
        id INTEGER NOT NULL PRIMARY KEY,
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
        id INTEGER NOT NULL PRIMARY KEY,
		user_id INTEGER NOT NULL,
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
	stmt := `
    CREATE TABLE IF NOT EXISTS tag (
        id INTEGER NOT NULL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		tag_name VARCHAR NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id)
    );
    `

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func createTodoTagTable(db *sql.DB) {
	stmt := `
    CREATE TABLE IF NOT EXISTS todo_tag (
        id INTEGER NOT NULL PRIMARY KEY,
		todo_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL,
		FOREIGN KEY(todo_id) REFERENCES todo(id)
		FOREIGN KEY(tag_id) REFERENCES tag(id)
    );
    `

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
