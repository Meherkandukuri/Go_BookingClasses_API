package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MeherKandukuri/studioClasses_API/helpers"
	"github.com/MeherKandukuri/studioClasses_API/models"
	"github.com/MeherKandukuri/studioClasses_API/repository"
)

// struct to hold payload from postrequest for creating class
type CreateClassRequest struct {
	ClassName string `json:"class_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

// struct to hold payload from postrequest for creating Booking
type BookingRequest struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type ClassRepository struct {
	ClassRepo repository.DBClassRepository
}

// Repo the repository used by the handlers
var ClassHandlerRepo *ClassRepository

func NewClassRepository(repo repository.DBClassRepository) *ClassRepository {
	return &ClassRepository{ClassRepo: repo}
}

type BookingRepository struct {
	bookingRepo repository.DBBookingRepository
	classRepo   ClassRepository
}

var BookingHandlerRepo *BookingRepository

func NewBookingRepository(bookingRepo repository.DBBookingRepository, classRepo ClassRepository) *BookingRepository {
	return &BookingRepository{bookingRepo: bookingRepo, classRepo: classRepo}
}

// Handler for postrequest for creating classes
func (h *ClassRepository) PostCreateClass(w http.ResponseWriter, r *http.Request) {
	// validating whether we got the right access method
	if !helpers.ValidateRequestMethod(w, r, http.MethodPost) {
		return
	}

	var req CreateClassRequest

	// loading the payload to variable for futher processing
	if !helpers.DecodeJSONPayload(w, r, &req) {
		return
	}

	// This function helps in validating the request.
	// we can check if the user did not enter any required field by comparing it to zero value of that particular field.
	// currently we validate only for zeros by adding "CheckZeroValue" to our checkstoBeDone slice
	checksToBeDone := []string{"checkZeroValue"}
	if !helpers.ValidateRequiredFields(w, req, checksToBeDone) {
		return
	}

	//parsing dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		http.Error(w, "Invalid endDate format", http.StatusBadRequest)
		return
	}

	// normalizing dates to a standard format
	startDate, endDate = helpers.NormalizeDate(startDate), helpers.NormalizeDate(endDate)

	// check if the dates entered are valid
	if startDate.After(endDate) {
		http.Error(w, "start date cannot be after end date", http.StatusBadRequest)
		return
	}

	class := models.Class{
		ClassName: req.ClassName,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  req.Capacity,
	}

	currentDate, err := h.ClassRepo.CreateClass(r.Context(), class)
	if err != nil {
		http.Error(w, fmt.Sprintf("Class already exists on %v", currentDate.Format("2006-01-02")), http.StatusConflict)
		return
	}
	// success message of creating a class
	message := fmt.Sprintf("created %s classes between %s and %s with Capacity: %d",
		class.ClassName, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"), class.Capacity)

	helpers.WriteJSONResponse(w, message, http.StatusCreated)

}

// Handler for Booking a class
func (h *BookingRepository) PostCreateBooking(w http.ResponseWriter, r *http.Request) {

	// validating whether we got the right access method
	if !helpers.ValidateRequestMethod(w, r, http.MethodPost) {
		return
	}

	// variable to hold the req data
	var reqBooking BookingRequest

	// loading the payload to variable for futher processing
	if !helpers.DecodeJSONPayload(w, r, &reqBooking) {
		return
	}

	// This function helps in validating the request.
	// we can check if the user did not enter any required field by comparing it to zero value of that particular field.
	// currently we validate only for zeroValues by adding "CheckZeroValue" to our checkstoBeDone slice
	checksToBeDone := []string{"checkZeroValue"}
	if !helpers.ValidateRequiredFields(w, reqBooking, checksToBeDone) {
		return
	}

	// parsing date
	date, err := time.Parse("2006-01-02", reqBooking.Date)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	// standardising the date for ease of comparision
	date = helpers.NormalizeDate(date)
	datestr := date.Format("2006-01-02")

	log.Println(datestr)
	// make sure we have a class on that date
	if _, found := h.classRepo.ClassRepo.GetClassByDate(r.Context(), date); !found {
		http.Error(w, "We don't have a class on this day", http.StatusBadRequest)
		return
	}

	// creating a struct for writing json response and storing to our in memory storage
	booking := models.Booking{
		Name: reqBooking.Name,
		Date: date,
	}
	log.Println(booking)
	// This check is done assuming there is only one name for one person.
	// later on We can achieve this functionality using unique user ID to make sure that all the bookings arent done by one person
	exists, _ := h.bookingRepo.BookingExists(r.Context(), date, booking.Name)
	if exists {
		http.Error(w, "You have already enrolled into class", http.StatusConflict)
		return
	}

	// appending to our bookings cache
	err = h.bookingRepo.CreateBooking(r.Context(), booking)
	if err != nil {
		http.Error(w, "Failed to Create Booking", http.StatusInternalServerError)
	}
	//writing to our response with a confirmation message
	message := fmt.Sprintf("%s has been enrolled for class on %s", booking.Name, datestr)
	helpers.WriteJSONResponse(w, message, http.StatusCreated)

}
