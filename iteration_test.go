package main

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleRepeat() {
	got := Repeat("i", 5)
	fmt.Println(got)
	// Output: iiiii
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	expected := "aaaaaa"

	if got != expected {
		t.Errorf("got %q but expected %q", got, expected)
	}
}

func TestCompare(t *testing.T) {
	stringA := "abc"
	stringB := "acd"

	// strings.Compare works lexicographically.
	if strings.Compare(stringA, stringB) != -1 {
		t.Errorf("Result after comparison should be -1.")
	}
}

func ExampleContains() {
	fmt.Println(strings.Contains("multiplicator", "Cat"))
	// fmt.Println(strings.Contains("multiplicator", "mul"))
	// Output: false
}
