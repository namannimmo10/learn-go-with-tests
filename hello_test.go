package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		// for helper functions, it's a good idea to accept a testing.TB
		// which acts as an interface that both *testing.T & *testing.B both
		// satisfy, so u can call helper functions from a test or a benchmark.
		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Naman!", "english")
		want := "Hello, Naman!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to hello world", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Nimmo!", "spanish")
		want := "Hola, Nimmo!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Nimmo!", "french")
		want := "Bonjour, Nimmo!"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloAgain(t *testing.T) {
	got := Hello("there", "english")
	want := "Hello, there!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
