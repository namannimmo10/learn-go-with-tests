package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(1 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	got := Racer(slowURL, fastURL)
	expected := fastURL

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}

	// Important to shut down the servers.
	slowServer.Close()
	fastServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
