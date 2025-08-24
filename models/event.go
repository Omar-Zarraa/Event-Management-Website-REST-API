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

func GetAllEvents() ([]Event, error) {
	selectQuery := `SELECT * FROM Events`

	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM Events WHERE ID = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {
	updateQuery :=
		`UPDATE Events
	SET Name = ?, Description = ?, Location = ?, Date = ?
	WHERE ID = ?`

	pQuery, err := db.DB.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer pQuery.Close()

	_, err = pQuery.Exec(event.Name, event.Description, event.Location, event.Date, event.ID)
	return err
}

func (event Event) DeleteEvent() error {
	delQuery := `DELETE FROM Events
	WHERE ID = ?`

	stmt, err := db.DB.Prepare(delQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
