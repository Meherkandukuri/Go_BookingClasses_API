package repository

import (
	"context"
	"time"

	"github.com/MeherKandukuri/studioClasses_API/models"
)

type DBClassRepository interface {
	CreateClass(ctx context.Context, class models.Class) (time.Time,error)
	GetClassByDate(ctx context.Context, date time.Time) (models.Class, bool)
}

type DBBookingRepository interface {
	CreateBooking(ctx context.Context, booking models.Booking) error
	BookingExists(ctx context.Context, date time.Time, name string) (bool, error)
}
