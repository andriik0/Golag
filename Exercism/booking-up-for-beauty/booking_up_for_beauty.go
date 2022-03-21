package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Bad date format!")
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Bad date format!")
	}
	return time.Since(t).Nanoseconds() > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Bad date format!")
	}
	return 12 <= t.Hour() && t.Hour() < 18 
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		panic("Bad date format!")
	}
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-09-15", time.Now().Year()))
	if err != nil {
		panic("Bad date format!")
	}
	return t

}
