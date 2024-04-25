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
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
 	user_id INTEGER
)`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Error creating tables")
	}
}

// func createTables() {
