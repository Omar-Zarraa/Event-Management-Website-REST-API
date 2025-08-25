//Package models handles the logic and structure of the events and users.
package models

import (
	"time"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/db"
)

//Event specifies the structure of the events to be used in the database.
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int64
}

//events is an empty Event
var events []Event = []Event{}

//Save inserts into the database the event that was used to call this method, possibly returns an error.
func (event *Event) Save() error {
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

//GetAllEvents returns all the events in the table as an Event slice, and possibly returns an error.
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

//GetEventByID takes the id of the desired event as a parameter and returns a pointer to the event, and possibly an error.
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

//UpdateEvent updates the event that was used to call this method in the database, and possibly returns an error.
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

//DeleteEvent deletes the event that was used to call this method from the database, and possibly returns an error.
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

//Register takes the userId as a parameter and inserts it into the Registrations table along with the eventId of the event used to call this method, possibly returns an error.
func (event *Event) Register(userId int64) error {
	query := "INSERT INTO Registrations(EventId,UserId) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}

//CancelRegistration takes the userId as a parameter and deletes the row containg the userId and the id of the event used to call this method from the Registrations table, possibly returns an error.
func (event Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM Registrations WHERE EventId = ? AND UserId = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}
