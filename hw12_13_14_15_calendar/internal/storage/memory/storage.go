package memorystorage

import (
	"sync"
	"time"

	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage"
	ltime "github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/time"
)

type Storage struct {
	mu     sync.RWMutex
	events map[int64]storage.Event
	idx    int64
}

func New() *Storage {
	return &Storage{
		events: map[int64]storage.Event{},
		idx:    1,
	}
}

func (s *Storage) Add(e storage.Event) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	e.ID = s.idx
	s.events[e.ID] = e

	s.idx++

	return e.ID, nil
}

func (s *Storage) Update(id int64, e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	event, ok := s.events[id]
	if !ok {
		return storage.ErrNoEvent
	}
	if event.UserId != e.UserId {
		return storage.ErrAccessDenied
	}

	e.ID = id
	s.events[id] = e

	return nil
}

func (s *Storage) Delete(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.events[id]
	if !ok {
		return storage.ErrNoEvent
	}
	delete(s.events, id)
	return nil
}

func (s *Storage) Get(id int64) (storage.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.events[id]
	if !ok {
		return storage.Event{}, storage.ErrNoEvent
	}

	return e, nil
}

func (s *Storage) ListForDate(inDate string) ([]storage.Event, error) {
	var result []storage.Event

	date, err := time.Parse(ltime.DateFormat, inDate)
	if err != nil {
		return nil, err
	}

	start, end := ltime.DayStartAndEnd(date)

	for _, event := range s.events {
		eDate, err := time.Parse(ltime.DateTimeFormat, event.DateStart)
		if err != nil {
			return nil, err
		}
		if ltime.DateInRange(eDate, start, end) {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *Storage) ListForWeek(inDate string) ([]storage.Event, error) {
	var result []storage.Event

	date, err := time.Parse(ltime.DateFormat, inDate)
	if err != nil {
		return nil, err
	}

	start, end := ltime.WeekStartAndEnd(date)

	for _, event := range s.events {
		eDate, err := time.Parse(ltime.DateTimeFormat, event.DateStart)
		if err != nil {
			return nil, err
		}
		if ltime.DateInRange(eDate, start, end) {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *Storage) ListForMonth(inDate string) ([]storage.Event, error) {
	var result []storage.Event

	date, err := time.Parse(ltime.DateFormat, inDate)
	if err != nil {
		return nil, err
	}

	start, end := ltime.MonthStartAndEnd(date)

	for _, event := range s.events {
		eDate, err := time.Parse(ltime.DateTimeFormat, event.DateStart)
		if err != nil {
			return nil, err
		}
		if ltime.DateInRange(eDate, start, end) {
			result = append(result, event)
		}
	}

	return result, nil
}
