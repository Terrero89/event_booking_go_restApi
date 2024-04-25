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

// fetching all the events in db
func GetAllEvents() ([]Event, error) {

		query := "SELECT * FROM events"
		rows, err := db.DB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	
		var events []Event
	
		for rows.Next() {
			var event Event
			err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	
			if err != nil {
				return nil, err
			}
	
			events = append(events, event)
		}
	
		return events, nil
	}

//Exec() is used when you are going to update/insert/delete data in db
//Query is used to extract/fetches data to see
