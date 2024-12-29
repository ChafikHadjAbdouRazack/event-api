package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (event *Event) Save() error {
	qurey := "INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)"
	statement, error := db.DB.Prepare(qurey)
	if error != nil {
		return error
	}
	defer statement.Close()

	result, error := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if error != nil {
		return error
	}
	id, error := result.LastInsertId()

	event.ID = id

	return error
}
func (event Event) Update() error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?"
	statement, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)

	return error
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	statement, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer statement.Close()
	_, error = statement.Exec(event.ID)

	if error != nil {
		return error
	}
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, error := db.DB.Query(query)
	if error != nil {
		return nil, error
	}
	var events []Event
	defer rows.Close()
	for rows.Next() {
		var event Event
		error := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if error != nil {
			return nil, error
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	error := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if error != nil {
		return nil, error
	}
	return &event, nil
}
