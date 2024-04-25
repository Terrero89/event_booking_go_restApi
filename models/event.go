package models

import (
	"restApi_go_event_booking/db"
	"time"
)

type Event struct {
	ID int64

	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
INSERT INTO events(id, name, description, location, dateTime, user_id)
VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() //closing db after prepare data to be sent to database
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	e.ID = id
	//later add it to a db
	//events = append(events, e)
	return err
}

func GetAllEvents() []Event {
	return events
}
