package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/MeherKandukuri/studioClasses_API/models"
)

// Testing Post Create class with a good request
func TestPostCreateClass_SuccessfulReq(t *testing.T) {

	// case 1: normal request sending right data and expecting status created
	reqBody := `{
	"class_name":"Pilates",
	"start_date":"2024-12-01",
	"end_date" : "2024-12-20",
	"capacity":10
	}`

	// creating a request with reqBody
	req:= httptest.NewRequest(http.MethodPost, "/classes", strings.NewReader(reqBody))

	// recorder for a responsewriter and creating an http handler from postCreateClass
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateClass)
	handler.ServeHTTP(rec, req)

	// checking for expected status code
	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// checking for expected message
	message := "created Pilates classes between 2024-12-01 and 2024-12-20 with Capacity: 10"
	expectedResponse := map[string]string{"message": message}

	var actualResponse map[string]string

	if err := json.Unmarshal(rec.Body.Bytes(), &actualResponse); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("handler returned unexpected body: got %v want %v", actualResponse, expectedResponse)
	}
}

// Test to check whether the function rejects invalid data
func TestPostCreateClasses_InvalidData(t *testing.T) {

	// Testing for invalid user data.
	requestBody := `{"class_name":"",
	"start_date":"2024-12-01",
	"end_date" : "2024-12-20",
	"capacity":10}` // we are sending a body with no name field

	req := httptest.NewRequest(http.MethodPost, "/classes", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateClass)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}

	expectedErrorMessage := "Missing or invalid value for field: ClassName"
	if rec.Body.String() != expectedErrorMessage+"\n" {
		t.Errorf("unexpected error message, got %s, expected %s", rec.Body.String(), expectedErrorMessage)
	}
}

// test to check whether function accepts invalidmethod request
func TestPostCreateClass_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/classes", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateClass)
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", rec.Code)
	}
}

// test to check whether the function accepts startdate after end date
func TestPostCreateBooking_WrongDates(t *testing.T) {

	// Testing for invalid user data.
	requestBody := `{"class_name":"Meher",
	"start_date":"2024-12-01",
	"end_date" : "2024-10-20",
	"capacity":10}` // we are sending a body with no name field

	req := httptest.NewRequest(http.MethodPost, "/classes", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateClass)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}

	expectedErrorMessage := "start date cannot be after end date"
	if rec.Body.String() != expectedErrorMessage+"\n" {
		t.Errorf("unexpected error message, got %s, expected %s", rec.Body.String(), expectedErrorMessage)
	}
}

// **************************************************
// started Tests for Booking Handler
// **************************************************


// Test to check wether the PostCreateBooking class as expected with a good request
func TestPostCreateBooking_SuccessfulReq(t *testing.T) {
	
	// Set up class for the test date and add it to cache
	classStorage = make(map[time.Time]models.Class)
	dateStr := "2024-10-02"
	date, _ := time.Parse("2006-01-02", dateStr)
	classStorage[date] = models.Class{
		ClassName: "Yoga",
		StartDate: date,
		EndDate:   date,
		Capacity:  20,
	}

	// set up a request Body
	requestBody := `{"name":"Meher",
				"date":"2024-10-02"}`

	// creating a request
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// creating and initializing a responseRecorder
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(PostCreateBooking)
	handler.ServeHTTP(rec, req)

	// check for status code
	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rec.Code)
	}


	// check for message
	expectedResponse := `{"message":"Meher has been enrolled for class on 2024-10-02"}`
	actualResponse := strings.TrimSpace(rec.Body.String())

	// Compare the two maps using reflect.DeepEqual
	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Errorf("expected message '%v', got '%v'", expectedResponse, actualResponse)
	}
}

// check if the function accepts required method
func TestPostCreateBooking_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/booking", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCreateClass)
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", rec.Code)
	}
}

// check if validation is working fine
func TestPostCreateBooking_InvalidDate(t *testing.T) {

	requestBody := `{"name":"Meher",
					"date":""}`
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(requestBody))

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(PostCreateBooking)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

// check if we function accepts booking when there is no class
func TestPostCreateBooking_NoClass(t *testing.T) {
	// Set up class for the test date and making sure that the dates are different as we 
	//dont want to have class on the booking day
	classStorage = make(map[time.Time]models.Class)
	dateStr := "2024-11-02"
	date, _ := time.Parse("2006-01-02", dateStr)
	classStorage[date] = models.Class{
		ClassName: "Yoga",
		StartDate: date,
		EndDate:   date,
		Capacity:  20,
	}

	requestBody := `{"name":"Meher",
				"date":"2024-10-02"}`
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(PostCreateBooking)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 201, got %d", rec.Code)
	}

	expectedResponse := `We don't have a class on this day`
	actualResponse := strings.TrimSpace(rec.Body.String())

	// Compare the reponses
	if expectedResponse != actualResponse {
		t.Errorf("expected message '%v', got '%v'", expectedResponse, actualResponse)
	}
}

// Testing for already existing booking
func TestPostCreateBooking_BookingExist(t *testing.T) {
	
	// Set up class for the test date
	classStorage = make(map[time.Time]models.Class)
	dateStr := "2024-11-02"
	date, _ := time.Parse("2006-01-02", dateStr)
	classStorage[date] = models.Class{
		ClassName: "Yoga",
		StartDate: date,
		EndDate:   date,
		Capacity:  20,
	}
	// create a bookings entry to test
	bookings = make(map[string][]string)
	bookings["2024-11-02"] = []string{"Meher"}
	
	//Creating a request body
	requestBody := `{"name":"Meher",
	"date":"2024-11-02"}`
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(PostCreateBooking)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Errorf("expected status 409, got %d", rec.Code)
	}

	expectedResponse := `You have already enrolled into class`
	actualResponse := strings.TrimSpace(rec.Body.String())

	// Compare the two maps using reflect.DeepEqual
	if expectedResponse != actualResponse {
		t.Errorf("expected message '%v', got '%v'", expectedResponse, actualResponse)
	}

}
