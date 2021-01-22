package main

// packages are way of grouping up related Go code together

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
