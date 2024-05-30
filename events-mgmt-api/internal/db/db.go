package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		panic("couldn't open the db")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := ` 
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	if _, err := DB.Exec(createEventsTable); err != nil {
		panic("couldn't create events table")
	}

	createUsersTable := ` 
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	if _, err := DB.Exec(createUsersTable); err != nil {
		panic("couldn't create users table")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	if _, err := DB.Exec(createRegistrationTable); err != nil {
		panic("couldn't create registrations table")
	}
}
