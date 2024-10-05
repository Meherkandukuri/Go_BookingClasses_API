package dbrepo

import (
	"time"

	"github.com/MeherKandukuri/studioClasses_API/models"
)

type inMemoryClassRepo struct {
	classStorage map[time.Time]models.Class
}

func NewinMemoryClassRepo() *inMemoryClassRepo {
	return &inMemoryClassRepo{
		classStorage: make(map[time.Time]models.Class),
	}
}

type inMemoryBookingRepo struct {
	bookings map[string][]string
}

func NewinMemoryBookingRepo() *inMemoryBookingRepo {
	return &inMemoryBookingRepo{
		bookings: make(map[string][]string),
	}
}
