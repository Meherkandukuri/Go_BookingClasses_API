package routes

import (
	"net/http"

	"github.com/MeherKandukuri/studioClasses_API/handlers"
	"github.com/go-chi/chi"
)

// Routes initializes and returns an HTTP handler with all the routes for the application.
func Routes() http.Handler {
	mux := chi.NewRouter()

	// -POST / classes: Handles the creating of class
	mux.Post("/classes", handlers.PostCreateClass)

	// -POSt /bookings: Handles the bookings for a class
	mux.Post("/bookings", handlers.PostCreateBooking)

	return mux
}
