package routes

import (
	"testing"

	"github.com/go-chi/chi"
)

// Check if the routes is returning the required mux
func TestRoutes(t *testing.T) {
	mux := Routes()

	switch v := mux.(type) {
		
	case *chi.Mux:
		// every thing is fine.
	default:
		t.Errorf("Type mismatch: Expected *chi.Mux, got %T", v)
	}
}