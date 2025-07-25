package models

import (
	"rest-api/db"
	"time"
)

type Events struct {
	ID         int64
	Title      string    `binding:"required"`
	UserId     int       `binding:"required"`
	Created_At time.Time `binding:"required"`
}

// var events = []Events{}

func (e Events) Save() {
	// events = append(events, e)
	query := `INSERT INTO events (title, userId, created_at) VALUES (?, ?, ?)`
	queryStatement, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer queryStatement.Close()
	result, err := queryStatement.Exec(e.Title, e.UserId, e.Created_At)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	e.ID = id

	if err != nil {
		panic(err)
	}

}

func GetAllEvents() ([]Events, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Events

	for rows.Next() {
		var event Events
		err := rows.Scan(&event.ID, &event.Title, &event.UserId, &event.Created_At)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEvent(id int64) (*Events, error) {
	query := "select * from events where id = ?"

	row := db.DB.QueryRow(query, id)
	var event Events

	err := row.Scan(&event.ID, &event.Title, &event.UserId, &event.Created_At)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Events) UpdateEvent(id int64) {
	query := `
	UPDATE events
	SET Title = ?, UserId = ?, Created_At = ?
	WHERE id = ?
	`
	queryStatement, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer queryStatement.Close()

	_, err = queryStatement.Exec(e.Title, e.UserId, e.Created_At, id)
	if err != nil {
		panic(err)
	}
}
