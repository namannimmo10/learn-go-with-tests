package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 6 numbers", func(t *testing.T) {
		numbers := [6]int{1, 2, 3, 5, 7, 0}
		got := Sum(numbers)
		expected := 18

		if got != expected {
			t.Errorf("got %d expected %d, sent %v", got, expected, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 3, 5}

		got := SumSlice(numbers)
		expected := 9

		if got != expected {
			t.Errorf("got %d expected %d, sent %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 4}, []int{4, -5})
	expected := []int{7, -1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}

	val1 := SumAll([]int{1, 0, 1})
	val2 := []int{2}

	if !reflect.DeepEqual(val1, val2) {
		t.Errorf("got %v, expected %v", val1, val2)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{2, 3, 5}, []int{7, 8, 9})
	expected := []int{8, 17}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
