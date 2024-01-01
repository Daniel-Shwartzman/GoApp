package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestHandleRequests(t *testing.T) {
	// Create a new HTTP request (a mock request for testing)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler from the handleRequests function
	handler := http.HandlerFunc(handleRequests)

	// Serve the HTTP request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response Heading
	expectedHeading := "<h1>test</h1>"
	if !strings.Contains(rr.Body.String(), expectedHeading) {
		t.Errorf("Handler returned body without expected heading: got %v want %v", rr.Body.String(), expectedHeading)
	}
}
