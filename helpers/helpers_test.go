package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestBooking struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

// Testing NormalizeDate function whether it is normalizing date or not
func TestNormalizeDate(t *testing.T) {
	// Input time with non-zero time components
	inputDate := time.Date(2024, time.October, 5, 15, 45, 30, 123456789, time.UTC)

	// Expected output with time normalized to midnight UTC
	expectedDate := time.Date(2024, time.October, 5, 0, 0, 0, 0, time.UTC)

	// Call the NormalizeDate function
	actualDate := NormalizeDate(inputDate)

	// Check if the normalized date matches the expected output
	if !actualDate.Equal(expectedDate) {
		t.Errorf("NormalizeDate failed: expected %v, got %v", expectedDate, actualDate)
	}
}

// Checking if the validateRequestMethod is working fine with valid request
func TestValidateRequestMethod_ValidRequestMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/classes", nil)
	rec := httptest.NewRecorder()

	expected := ValidateRequestMethod(rec, req, http.MethodPost)

	if !expected {
		t.Errorf("expected result to be true, got false")
	}
}

// we are sending an invalid request so that the valiation should fail in below case
func TestValidateRequestMethod_InvalidRequestMethod(t *testing.T) {
	// Create a request with a different method (GET instead of POST)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Call the function with a different expected method
	expected := ValidateRequestMethod(rec, req, http.MethodPost)

	// Check that the function returned false
	if expected {
		t.Errorf("expected result to be false, got true")
	}

	// Check that the response recorder captured the "method not allowed" error
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", rec.Code)
	}
}

// sending a acceptable payload to DecodeJsonPayload function
func TestDecodeJSONPayload_ValidPayload(t *testing.T) {

	var payload TestBooking

	// create a valid JSON
	requestBody := `{"name":"Meher",
				"date":"2024-10-02"}`

	// create a new Http request with the JSON payload
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call the function and check if true is returned as expected
	expected := DecodeJSONPayload(rec, req, &payload)

	if !expected {
		t.Errorf("expected success to be true, got false")
	}

	if payload.Name != "Meher" || payload.Date != "2024-10-02" {
		t.Errorf("unexpected decoded payload: got %+v", payload)
	}

}

// Here we test with an invalid payload
func TestDecodeJSONPayload_InValidPayload(t *testing.T) {
	var payload TestBooking

	// create an invalid JSON
	requestBody := `{"name":"Meher",
	"date":"2024-10-02"`

	// create a new Http request with the JSON payload
	req := httptest.NewRequest(http.MethodPost, "/classics", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call the function and check that it returns true
	expected := DecodeJSONPayload(rec, req, &payload)

	if expected {
		t.Errorf("expected should be false but got true")
	}
}

func TestWriteJSONResponse(t *testing.T) {
	// Set up the response recorder
	rec := httptest.NewRecorder()

	// Call the function with a sample message and status code
	WriteJSONResponse(rec, "Successfully written JSON object", http.StatusOK)

	// Check the content type
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected content type 'application/json', got '%s'", rec.Header().Get("Content-Type"))
	}

	// Check the status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code 200, got %d", rec.Code)
	}

	// Parse the response body
	var responseBody map[string]string
	err := json.NewDecoder(rec.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check the message in the response body
	expectedMessage := "Successfully written JSON object"
	if responseBody["message"] != expectedMessage {
		t.Errorf("expected message '%s', got '%s'", expectedMessage, responseBody["message"])
	}
}

// Its a check for ValidateRequiredFields function is running properly with a valid response.
func TestValidateRequiredFields(t *testing.T) {
	checklist := []string{"checkZeroValue"}
	rec := httptest.NewRecorder()
	reqPayload := TestBooking{
		Name: "Meher",
		Date: "2024-06-07",
	}

	expected := ValidateRequiredFields(rec, reqPayload, checklist)

	if !expected {
		t.Errorf("expected should be true but got false")
	}

}
