package config

import (
	"time"
)

// MonthPeriod returns the first and last day of the month
func MonthPeriod(t time.Time) (time.Time, time.Time) {
	year, month, _ := t.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return start, end
}

// printDate prints a time as d/m/y
func printDate(t time.Time) string {
	return t.Format("02/01/2006")
}
