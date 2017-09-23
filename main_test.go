package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandle tests the generic request handler
func TestHandle(t *testing.T) {
	var (
		req  = `{"scooters": [11, 15, 13], "C": 10, "P": 5}`
		want = `{"fleet_engineers":7}`
		buf  bytes.Buffer
	)

	if err := handle(&buf, strings.NewReader(req)); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got := strings.Trim(buf.String(), " \n"); got != want {
		t.Errorf("unexpected response written: got %s, want %s", got, want)
	}
}

// TestHandleHTTP tests the HTTP handler
func TestHandleHTTP(t *testing.T) {
	var (
		body    = `{"scooters": [11, 15, 13], "C": 10, "P": 5}`
		handler = handleHTTP()
		rec     = httptest.NewRecorder()
		want    = `{"fleet_engineers":7}`
	)

	req, err := http.NewRequest("POST", "/", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rec, req)

	if got, want := rec.Code, http.StatusOK; got != want {
		t.Errorf("unexpected status code: got %v want %v", got, want)
	}

	if got := strings.Trim(rec.Body.String(), " \n"); got != want {
		t.Errorf("unexpected response body: got %v want %v", got, want)
	}
}
