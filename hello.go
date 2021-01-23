package main

// packages are way of grouping up related Go code together

import (
	"fmt"
	"strings"
)

const MAX = 4000000
const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Public function; so starts with an upper case.
func Hello(name string, lang string) string {
	if name == "" {
		name = "World!"
	}

	if name == "there" {
		name = "there!"
	}

	mainPrefix := englishHelloPrefix

	switch lang {
	case french:
		mainPrefix = frenchHelloPrefix
	case spanish:
		mainPrefix = spanishHelloPrefix
	}

	return mainPrefix + name
}

func Mul(x int, y float64) float64 {
	return float64(x) * y
}

func Repeat(char string, repeatCount int) string {
	var repeatedString string
	for i := 0; i < repeatCount; i++ {
		repeatedString += char
	}
	return repeatedString
}

// Prints the sum of even-valued fibonacci terms under 4 million
func SumEvenFibonacci() int {
	a := 1
	b := 2
	sum := 2
	next := a + b

	for next < MAX {
		if (next & 1) == 0 {
			sum += next
		}
		a = b
		b = next
		next = a + b
	}

	return sum
}

func main() {
	fmt.Println(Hello("world", "english"))
	fmt.Println(Mul(5, 6.2), Mul(2, 3.3))
	fmt.Println(Repeat("p", 0) == "")
	fmt.Println(strings.Contains("substr", "Str"))
	fmt.Println(SumEvenFibonacci())
}
