package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nimmo")

	got := buffer.String()
	// String returns the contents of the unread portion
	// of the buffer as a string.

	expected := "Hello, Nimmo"
	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
