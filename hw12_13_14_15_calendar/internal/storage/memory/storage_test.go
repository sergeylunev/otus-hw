package memorystorage

import (
	"testing"

	"github.com/sergeylunev/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		s := New()

		e := storage.Event{
			Title:       "test",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is just a test",
			UserId:      100500,
		}

		id, err := s.Add(e)

		require.Nil(t, err)
		require.EqualValues(t, 1, id)
	})
	t.Run("get", func(t *testing.T) {
		s := New()

		e := storage.Event{
			Title:       "test",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is just a test",
			UserId:      100500,
		}

		id, _ := s.Add(e)
		ge, err := s.Get(id)

		require.Nil(t, err)
		require.Equal(t, "test", ge.Title)
		require.Equal(t, "this is just a test", ge.Description)
		require.EqualValues(t, 1, ge.ID)
	})
	t.Run("update", func(t *testing.T) {
		s := New()

		e := storage.Event{
			Title:       "test",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is just a test",
			UserId:      100500,
		}

		id, _ := s.Add(e)

		ue := storage.Event{
			Title:       "test upd",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is just a test and its updated",
			UserId:      100500,
		}

		err := s.Update(id, ue)

		require.Nil(t, err)

		ge, _ := s.Get(id)

		require.Equal(t, "test upd", ge.Title)
		require.Equal(t, "this is just a test and its updated", ge.Description)
		require.EqualValues(t, 1, ge.ID)
	})
	t.Run("delete", func(t *testing.T) {
		s := New()

		e := storage.Event{
			Title:       "test",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is just a test",
			UserId:      100500,
		}

		id, _ := s.Add(e)
		err := s.Delete(id)
		require.Nil(t, err)

		_, err = s.Get(id)

		require.ErrorIs(t, err, storage.ErrNoEvent)
	})

	t.Run("get day events", func(t *testing.T) {
		s := New()

		e1 := storage.Event{
			Title:       "true",
			DateStart:   "30.12.2021 10:00:00",
			DateEnd:     "30.12.2021 11:30:00",
			Description: "this event we looking for",
			UserId:      100500,
		}
		e2 := storage.Event{
			Title:       "false",
			DateStart:   "31.12.2021 10:00:00",
			DateEnd:     "31.12.2021 11:30:00",
			Description: "this is wrong event",
			UserId:      100500,
		}
		e3 := storage.Event{
			Title:       "false",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is wrong event",
			UserId:      100500,
		}
		s.Add(e1)
		s.Add(e2)
		s.Add(e3)

		events, err := s.ListForDate("30.12.2021")

		require.Nil(t, err)
		require.Len(t, events, 1)
	})

	t.Run("get week events", func(t *testing.T) {
		s := New()

		e1 := storage.Event{
			Title:       "true",
			DateStart:   "30.12.2021 10:00:00",
			DateEnd:     "30.12.2021 11:30:00",
			Description: "this event we looking for",
			UserId:      100500,
		}
		e2 := storage.Event{
			Title:       "true",
			DateStart:   "31.12.2021 22:00:00",
			DateEnd:     "01.01.2022 07:30:00",
			Description: "this is event we looking for",
			UserId:      100500,
		}
		e3 := storage.Event{
			Title:       "false",
			DateStart:   "22.12.2021 22:00:00",
			DateEnd:     "22.12.2021 23:30:00",
			Description: "this is wrong event",
			UserId:      100500,
		}
		s.Add(e1)
		s.Add(e2)
		s.Add(e3)

		events, err := s.ListForWeek("30.12.2021")

		require.Nil(t, err)
		require.Len(t, events, 2)
	})
	t.Run("get month events", func(t *testing.T) {
		s := New()

		e1 := storage.Event{
			Title:       "true",
			DateStart:   "30.12.2021 10:00:00",
			DateEnd:     "30.12.2021 11:30:00",
			Description: "this event we looking for",
			UserId:      100500,
		}
		e2 := storage.Event{
			Title:       "true",
			DateStart:   "22.12.2021 22:00:00",
			DateEnd:     "22.12.2021 23:30:00",
			Description: "this is event we looking for",
			UserId:      100500,
		}
		e3 := storage.Event{
			Title:       "false",
			DateStart:   "22.11.2021 22:00:00",
			DateEnd:     "22.11.2021 23:30:00",
			Description: "this is wrong event",
			UserId:      100500,
		}
		s.Add(e1)
		s.Add(e2)
		s.Add(e3)

		events, err := s.ListForMonth("30.12.2021")

		require.Nil(t, err)
		require.Len(t, events, 2)
	})
}
