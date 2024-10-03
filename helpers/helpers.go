package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// Helper function is used to standardize the time component and maintain consistency througout the code
func NormalizeDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

// Function used to validate the request method is of acceptable method
func ValidateRequestMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, fmt.Sprintf("Invalid request method: expected %s", method), http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// Function used to DecodeJSONOayLoad to a variable
func DecodeJSONPayload(w http.ResponseWriter, r *http.Request, reqVariable any) bool {
	if err := json.NewDecoder(r.Body).Decode(reqVariable); err != nil {
		http.Error(w, "Unable to decode the request body payload", http.StatusBadRequest)
		return false
	}
	return true
}

// helper function to write jsonresponse to Response writer with a required status code
func WriteJSONResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": message}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// helper function can be used to validate a slice of checks.
// acceptable values in checks slice are :
// 1) "checkZeroValue": used to check if user didnt fill the fields or may be few fileds are missing.
func ValidateRequiredFields(w http.ResponseWriter, reqPayload any, checks []string) bool {
	for _, checkType := range checks {

		switch checkType {

		case "checkZeroValue":
			t := reflect.TypeOf(reqPayload)

			// accepts only if the type is struct
			if t.Kind() == reflect.Struct {
				val := reflect.ValueOf(reqPayload)

				for i := 0; i < val.NumField(); i++ {
					fieldValue := val.Field(i)

					if isZero(fieldValue) {
						fieldName := t.Field(i).Name
						http.Error(w, "Missing or invalid value for field: "+fieldName, http.StatusBadRequest)
						return false
					}

				}
				return true
			}
		}

	}
	return false
}

// As we cannot compare the zero value of a data type with reflect value directly,
// This function helps to find if the value of a field is its zerovalue
func isZero(v reflect.Value) bool {
	// check if the value is zero for its type
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
