package main

func Sum(numbers [6]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Variadic function. Another example is `Println`.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}

// Returns a slice which contains the sum of "tails" of each slice.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		numbers = numbers[1:] // tail of each slice.
		sums = append(sums, SumSlice(numbers))
	}

	return sums
}
