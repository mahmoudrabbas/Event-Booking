package models

import (
	"errors"
	"fmt"
	"time"

	"example.com/events/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

// var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) values (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("Error In Preparing query")
	}

	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return errors.New("Error In Executing the query")
	}

	id, err := res.LastInsertId()

	if err != nil {
		return errors.New("Error In getting lastInsertId")

	}

	e.Id = id
	return nil
}

func GetEvents() ([]Event, error) {
	fmt.Println("s")
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetSingleEvent(id int64) (*Event, error) {
	fmt.Println("Gddd1")

	query := `SELECT * FROM events where id = ?`

	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	fmt.Println("end of the get func")
	return &event, nil
}

func (e Event) UpdateEvent() error {

	fmt.Println("1")
	query := `UPDATE events SET
	name = ?,
	description = ?,
	location = ?,
	dateTime = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("Error Preparing the query")
	}

	defer stmt.Close()
	// var event Event

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	if err != nil {
		return errors.New("Error Updating the Event.")
	}

	return err

}

func (e *Event) DeleteEvent() error {
	query := "DELETE FROM events where id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("Couldnt Prepare Query.")
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	return err

}
