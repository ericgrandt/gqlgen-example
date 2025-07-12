package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetOrCreateDatabase(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)

	return db
}

func createTable(db *sql.DB) {
	stmt := `
    CREATE TABLE IF NOT EXISTS todo (
        id int NOT NULL PRIMARY KEY,
		value varchar NOT NULL
    );
    `

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}
