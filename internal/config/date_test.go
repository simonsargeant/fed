package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate_MonthPeriod(t *testing.T) {
	for s, tc := range map[string]struct {
		input time.Time
		start time.Time
		end   time.Time
	}{
		"30 day month": {
			input: time.Date(2020, time.April, 5, 1, 1, 1, 1, time.UTC),
			start: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.April, 30, 0, 0, 0, 0, time.UTC),
		},
		"31 day month": {
			input: time.Date(2020, time.March, 1, 1, 1, 1, 1, time.UTC),
			start: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.March, 31, 0, 0, 0, 0, time.UTC),
		},
		"Last day of the month": {
			input: time.Date(2020, time.April, 30, 1, 1, 1, 1, time.UTC),
			start: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.April, 30, 0, 0, 0, 0, time.UTC),
		},
		"February": {
			input: time.Date(2021, time.February, 9, 1, 1, 1, 1, time.UTC),
			start: time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2021, time.February, 28, 0, 0, 0, 0, time.UTC),
		},
		"February on a leap year": {
			input: time.Date(2020, time.February, 20, 1, 1, 1, 1, time.UTC),
			start: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC),
		},
	} {
		tc := tc
		t.Run(s, func(t *testing.T) {
			t.Parallel()

			start, end := MonthPeriod(tc.input)

			assert.Equal(t, tc.start, start)
			assert.Equal(t, tc.end, end)
		})
	}
}

func TestDate_PrintDate(t *testing.T) {
	for s, tc := range map[string]struct {
		input time.Time
		res   string
	}{
		"Some date": {
			input: time.Date(2020, time.April, 5, 1, 1, 1, 1, time.UTC),
			res:   "05/04/2020",
		},
	} {
		tc := tc
		t.Run(s, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.res, printDate(tc.input))
		})
	}
}
