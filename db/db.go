package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Email TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS Events(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Description TEXT NOT NULL,
		Location TEXT NOT NULL,
		Date DATETIME NOT NULL,
		UserId INTEGER,
		FOREIGN KEY(UserId) REFERENCES Users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}
