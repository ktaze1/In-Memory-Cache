package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(notAllowedHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFlushHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/flush", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(flushHandler)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// TODO: Add more tests for GET and PUT
