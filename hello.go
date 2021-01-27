package main

// packages are way of grouping up related Go code together

import (
	"fmt"
	"math"
	"strings"
)

const MAX = 4000000
const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) area() float64 {
	return 0.5 * t.Base * t.Height
}

type Shape interface {
	// Interfaces allow us to make functions that can
	// be used with different types.
	area() float64
}

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

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.height * rectangle.width
}

func Repeat(char string, repeatCount int) string {
	var repeatedString string
	for i := 0; i < repeatCount; i++ {
		repeatedString += char
	}
	return repeatedString
}

func Slices() {
	s := make([]byte, 5)
	var nilSlice []int

	fmt.Println(s, len(s), cap(s))
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))

	// Internally slices work this way!
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}

	for i := range d {
		fmt.Println(d[i])
	}
}

// Prints the sum of even-valued fibonacci terms under 4 million
func SumEvenFibonacci() int {
	a, b, sum := 1, 2, 2
	next := a + b

	for next < MAX {
		if (next & 1) == 0 {
			sum += next
		}
		a, b = b, next
		next = a + b
	}

	return sum
}

func IsPalindrome(num int) bool {
	reversed, originalNum := 0, num

	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}

	if originalNum == reversed {
		return true
	} else {
		return false
	}
}

func LargestPalindrome() int {
	result, a := 0, 999

	for a >= 100 {
		b := 999
		for b >= a {
			prod := a * b
			if IsPalindrome(prod) {
				if prod > result {
					result = prod
				}
			}
			b--
		}
		a--
	}

	return result
}

func LargestPrimeFactor(num int) int {
	j := 2

	for (num % j) != 0 {
		j += 1
	}

	if (num % 2) == 0 {
		if num == 2 {
			return 2
		} else {
			return LargestPrimeFactor(num / 2)
		}
	} else if num/j == 1 {
		return num
	} else {
		return LargestPrimeFactor(num / j)
	}
}

func main() {
	// map ternary!
	c := map[bool]int{true: 1, false: 0}[5 < 4]
	fmt.Println(c)

	Slices()

	// An array var denotes the entire array,
	// it is not a pointer to the first array elem.
	a := [...]string{"english", "spanish", "french"}
	fmt.Println(len(a), cap(a))

	fmt.Println(Hello("world", "english"))
	fmt.Println(Mul(5, 6.2), Mul(2, 3.3))
	fmt.Println(Repeat("p", 0) == "")
	fmt.Println(strings.Contains("substr", "Str"))
	fmt.Println(SumEvenFibonacci())
	fmt.Println(LargestPrimeFactor(13195))        // 29
	fmt.Println(LargestPrimeFactor(600851475143)) // 6857
	fmt.Println(LargestPalindrome())              // 906609
}
