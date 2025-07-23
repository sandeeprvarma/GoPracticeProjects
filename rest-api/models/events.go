package models

import "time"

type Events struct {
	ID         int
	Title      string `binding:"required"`
	UserId     int    `binding:"required"`
	Created_At time.Time
}

var events = []Events{}

func (e Events) Save() {
	events = append(events, e)
}

func GetAllEvents() []Events {
	return events
}
