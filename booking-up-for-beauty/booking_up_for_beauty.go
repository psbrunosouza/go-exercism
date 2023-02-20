package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var layout string = "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var layout string = "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return time.Now().Compare(parsedDate) == 1
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var layout string = "Monday, January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate.Hour() >= 12 && parsedDate.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var time time.Time = Schedule(date)
	var layout string = "Monday, January 2, 2006, at 15:04"
	return fmt.Sprintf("You have an appointment on %s.", time.Format(layout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
