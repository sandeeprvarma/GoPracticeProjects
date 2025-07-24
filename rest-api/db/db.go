package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func DbInit() {
	var err error
	DB, err = sql.Open("sqlite3", "restapi.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		userId INTEGER NOT NULL,
		created_at DATETIME NOT NULL
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
