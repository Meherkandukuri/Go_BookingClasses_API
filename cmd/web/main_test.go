package main

import "testing"

func TestRun(t *testing.T) {
	server := run()

	// Check we are running on required port:
	if server.Addr != ":8080" {
		t.Errorf("expected server to run on address 8080 but running on %s", server.Addr)
	}

	// check we have a handler
	if server.Handler == nil {
		t.Errorf("expected a server with handler but got nil")
	}

}
