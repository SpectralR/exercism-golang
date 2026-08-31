package advanced

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	fmt.Println(date)
	//t, err := time.Parse(date, date)
	t, err := time.Parse("", "7/13/2020 20:32:00")
	fmt.Println(err)
	fmt.Println(t)
	if err != nil {
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointment, err := time.Parse("2019-07-25 13:45:00", date)
	if err != nil {
	}

	return appointment.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment, err := time.Parse("13:45:00", date)
	start, err := time.Parse("12:00:00", "12:00:00")
	end, err := time.Parse("18:00:00", "18:00:00")
	if err != nil {
	}

	return appointment.After(start) && appointment.Before(end)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed, err := time.Parse("Thursday, July 25, 2019, at 13:45", date)
	if err != nil {
	}

	return "You have an appointment on " + parsed.String()
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	bday, err := time.Parse("2020-09-15", "15-09")
	if err != nil {
	}
	return bday
}
