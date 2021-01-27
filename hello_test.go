package main

import (
	"fmt"
	"testing"
)

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

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 12.0}
	got := Perimeter(rectangle)
	expected := 44.0

	if got != expected {
		t.Errorf("got %.3f, expected %.3f", got, expected)
	}
}

func TestArea(t *testing.T) {
	// Table-based tests are a great fit if we wish to test
	// various implementations of an interface, or if the data
	// passed in to a func has lots of different requirements.
	areaTests := []struct {
		name  string
		shape Shape
		Area  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, Area: 72.0},
		{name: "Circle", shape: Circle{10}, Area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, Area: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.area()
			if got != tt.Area {
				t.Errorf("%#v: got %g, expected %g", tt.name, got, tt.Area)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(56765) {
		t.Errorf("The number passed in is a palindrome.")
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome(56765))
	// Output: true
}
