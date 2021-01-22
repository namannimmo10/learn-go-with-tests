package main

import (
	"fmt"
	"testing"
)

func TestMultiplicator(t *testing.T) {
	prod := Mul(3, 4)
	expected := float64(12)

	if prod != expected {
		t.Errorf("expected %f but got %f", expected, prod)
	}
}

func ExampleMul() {
	prod := Mul(2, 3.3)
	fmt.Println(prod)
	// Output: 6.6
}
