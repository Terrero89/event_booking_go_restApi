package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// global var to point to sql struct
var DB *sql.DB

// database initialization.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	//check connnection, if not, panic.
	if err != nil {
		panic("Connection not possible to DB")
	}
	//determines how many connections we can have while sql is open
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
