package models

import (
	"time"

	"github.com/Omar-Zarraa/REST-API/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (event Event) Save() error {
	insertQuery := `INSERT INTO Events (Name, Description,Location,Date,UserId)
	VALUES (?,?,?,?,?)`

	pQuery, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer pQuery.Close()

	result, err := pQuery.Exec(event.Name, event.Description, event.Location, event.Date, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetAllEvents() ([]Event,error) {
	selectQuery := `SELECT * FROM Events`

	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil,err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}

	return events,nil
}

//func AddEvent(id, userId int, name, description, location string, date time.Time)
