package dbrepo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MeherKandukuri/studioClasses_API/models"
)

func (m *inMemoryClassRepo) CreateClass(ctx context.Context, class models.Class) (time.Time, error) {
	currentDate := class.StartDate

	// If there is a class on that Day, we cannot create one more as we should have only one class per day
	for !currentDate.After(class.EndDate) {
		if _, exists := m.classStorage[currentDate]; exists {
			return currentDate, fmt.Errorf("class already exists on %v", currentDate.Format("2006-01-02"))
		}
		m.classStorage[currentDate] = class
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return time.Time{}, nil
}

func (m *inMemoryClassRepo) GetClassByDate(ctx context.Context, date time.Time) (models.Class, bool) {

	class, ok := m.classStorage[date]
	return class, ok

}

func (m *inMemoryBookingRepo) CreateBooking(ctx context.Context, booking models.Booking) error {
	dateStr := booking.Date.Format("2006-01-02")
	m.bookings[dateStr] = append(m.bookings[dateStr], booking.Name)
	return nil
}

func (m *inMemoryBookingRepo) BookingExists(ctx context.Context, date time.Time, name string) (bool, error) {
	dateStr := date.Format("2006-01-02")

	namesInClass := m.bookings[dateStr]
	username := strings.ToLower(strings.TrimSpace(name))

	for _, name := range namesInClass {

		if strings.ToLower(strings.TrimSpace(name)) == username {
			return true, nil
		}
	}

	return false, nil
}
