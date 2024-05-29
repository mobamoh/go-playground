package event

import (
	"time"

	"github.com/mobamoh/go-playground/events-mgmt-api/internal/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      string
}

func (event Event) Save() error {
	query := `INSERT INTO events(name, description, location, date_time, user_id)
	VALUES(?, ?, ?, ?, ?)	
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	event.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	var events []Event
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		if err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdatEvent() error {

	query := `UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return err
}
