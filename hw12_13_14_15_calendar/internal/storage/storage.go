package storage

import (
	"errors"
)

var (
	ErrNoEvent          = errors.New("no such event")
	ErrEventsWithSameId = errors.New("events with same id")
	ErrAccessDenied     = errors.New("access for event denied")
)

type Storage interface {
	Add(event Event) (int64, error)
	Update(id int64, event Event) error
	Delete(id int64) error
	Get(id int64) (Event, error)
	ListForDate(date string) ([]Event, error)
	ListForWeek(date string) ([]Event, error)
	ListForMonth(date string) ([]Event, error)
}

type StorageConf struct {
	Type string
	User string
	Pass string
	Port string
	Host string
	Name string
}
