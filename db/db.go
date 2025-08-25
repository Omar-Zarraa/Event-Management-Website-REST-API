// Package db is the package handling the database.
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database variable.
var DB *sql.DB

// InitDB initializes the database of type sqlite3 and calls on the createTables function.
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

// createTables initializes the three tables used; Users(ID, Email, Password), Events(ID, Name, Description, Location, Date, UserId), and Registrations(ID, EventId, UserId).
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

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS Registrations (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		EventId INTEGER,
		UserId INTEGER,
		FOREIGN KEY(EventId) REFERENCES Events(ID),
		FOREIGN KEY(UserId) REFERENCES Users(ID)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registrations table")
	}
}
