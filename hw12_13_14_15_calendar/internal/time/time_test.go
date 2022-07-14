package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	dateFormat     = "02.01.2006"
	datetimeFormat = "02.01.2006 15:04"
)

func TestTime(t *testing.T) {
	t.Run("test day defenition", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "31.12.2021")

		start, end := DayStartAndEnd(day)

		require.Equal(t, "31.12.2021 00:00", start.Format(datetimeFormat))
		require.Equal(t, "31.12.2021 23:59", end.Format(datetimeFormat))
	})

	t.Run("test week start and end", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "30.06.2022")
		start, end := WeekStartAndEnd(day)

		require.Equal(t, "27.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "03.07.2022 23:59", end.Format(datetimeFormat))
	})
	t.Run("test week start and end sunday", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "03.07.2022")
		start, end := WeekStartAndEnd(day)

		require.Equal(t, "27.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "03.07.2022 23:59", end.Format(datetimeFormat))
	})
	t.Run("test week start and end monday", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "27.06.2022")
		start, end := WeekStartAndEnd(day)

		require.Equal(t, "27.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "03.07.2022 23:59", end.Format(datetimeFormat))
	})

	t.Run("test month start and end", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "10.06.2022")
		start, end := MonthStartAndEnd(day)

		require.Equal(t, "01.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "30.06.2022 23:59", end.Format(datetimeFormat))
	})
	t.Run("test month start and end first day", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "01.06.2022")
		start, end := MonthStartAndEnd(day)

		require.Equal(t, "01.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "30.06.2022 23:59", end.Format(datetimeFormat))
	})
	t.Run("test month start and end lastday", func(t *testing.T) {
		day, _ := time.Parse(dateFormat, "30.06.2022")
		start, end := MonthStartAndEnd(day)

		require.Equal(t, "01.06.2022 00:00", start.Format(datetimeFormat))
		require.Equal(t, "30.06.2022 23:59", end.Format(datetimeFormat))
	})
}
