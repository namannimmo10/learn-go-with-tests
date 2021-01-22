package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World!")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}

// func TestHelloAgain(t *testing.T) {
// 	got := Hello("world")
// 	want := "Jello World!"

// 	if got != want {
// 		t.Errorf("got %q but wanted %q", got, want)
// 	}
// }
