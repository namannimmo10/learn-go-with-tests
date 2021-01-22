package main

// packages are way of grouping up related Go code together

import "fmt"

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

func main() {
	fmt.Println(Hello("world", "english"))
	fmt.Println(Mul(5, 6.2), Mul(2, 3.3))
}
