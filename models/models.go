package models

import "time"

// used to store class data
type Class struct {
	ClassName string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

// used to store booking data
type Booking struct {
	Name string
	Date time.Time
}
