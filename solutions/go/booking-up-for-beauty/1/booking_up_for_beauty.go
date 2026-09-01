package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointment, _ := time.Parse("January 2, 2006 15:04:05", date)
	return appointment.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return 12 <= appointment.Hour() && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	appointment, _ := time.Parse(layout, date)
	day := appointment.Format("Monday, January 2, 2006")
	timeOfDay := appointment.Format("15:04")
	return "You have an appointment on " + day + ", at " + timeOfDay + "."

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	date := "15-09-2026"
	format := "2-01-2006"
	anniversary, _ := time.Parse(format, date)
	return anniversary
}
