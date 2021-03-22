package main

// packages are way of grouping up related Go code together

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	countdownStart     = 3
	MAX                = 4000000
	finalWord          = "Go!"
	spanish            = "spanish"
	french             = "french"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

var z complex128 = 5 + 6i

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

type result struct {
	string
	bool
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human      // embedded field
	speciality string
	phone      string
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

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
	// `fmt.Fprintf` is just like `fmt.Printf`;
	// But takes a Writer to send the string to,
	// whereas `fmt.Printf` defaults to stdout.
}

// MyGreeterHandler says "Hello, world" over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world") // `http.ResponseWriter` also implements `io.Writer`.
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

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
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

func SwapCase(r rune) rune {
	switch {
	case r >= 97 && r <= 122:
		return r - 32
	case r >= 65 && r <= 90:
		return r + 32
	default:
		return r
	}
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {

	var result strings.Builder
	// A `Builder` is used to efficiently build a string
	// using Write methods. It minimizes memory copying.

	for _, romanNumeral := range allRomanNumerals {
		for arabic >= romanNumeral.Value {
			result.WriteString(romanNumeral.Symbol)
			arabic -= romanNumeral.Value
		}
	}

	return result.String()
}

type WebsiteChecker func(string) bool

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// `CheckWebsites` takes a `WebsiteChecker` and a slice of urls and returns
// a map of urls to the result of checking each url with `WebsiteChecker`.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// Creates a new file using `os.File`.
func createFile(p string) *os.File {
	fmt.Println("Creating a file")
	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}

	return f
}

// Writes to a new file using `Fprintln`.
func writeFile(f *os.File) {
	fmt.Println("Writing to the file.")
	fmt.Fprintln(f, "this is a new file!")
}

func closeFile(f *os.File) {
	fmt.Println("Closing the file")
	err := f.Close()

	if err != nil {
		fmt.Println("Some issue with closing the file.")
		os.Exit(1)
	}
}

func main() {
	Countdown(os.Stdout)
	fmt.Println()

	// map ternary!
	c := map[bool]int{true: 1, false: 0}[5 < 4]
	fmt.Println(c)

	Slices()

	// An array var denotes the entire array,
	// it is not a pointer to the first array elem.
	a := [...]string{"english", "spanish", "french"}
	fmt.Println(len(a), cap(a))

	for val := range a {
		fmt.Printf("%d\n", val)
	}

	k := 1
	doubleArray := [3][2]int{{1, 2}, {0, 1}, {0, 0}}

Here:
	k++
	fmt.Println(k)
	if k != 5 {
		goto Here
	}

	fmt.Println(rune('g'), rune('o'), rune('G'))
	fmt.Println(Hello("world", "english"))
	fmt.Println(Mul(5, 6.2), Mul(2, 3.3))
	fmt.Println(Repeat("p", 0) == "")
	fmt.Println(strings.Contains("substr", "Str"))
	fmt.Println(SumEvenFibonacci())
	fmt.Println(LargestPrimeFactor(13195))        // 29
	fmt.Println(LargestPrimeFactor(600851475143)) // 6857
	fmt.Println(LargestPalindrome())              // 906609
	fmt.Printf("value of the complex no. is: %v\n", z)
	fmt.Println(doubleArray)
	Greet(os.Stdout, "Nimo\n")

	// err := http.ListenAndServe(":5000", http.HandleFunc(MyGreeterHandler))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's phone no. is:", Bob.phone) // field overloading
	fmt.Println("Bob's personal phone no. is", Bob.Human.phone)

	kvp := map[string]string{"a": "val1", "b": "val2"}
	for key, val := range kvp {
		fmt.Printf("%s -> %s\n", key, val)
	}

	for key := range kvp {
		fmt.Printf("%s\n", key)
	}

	resp, err := http.Get("https://namannimmo10.github.io/emerald//week_12")
	if err != nil {
		log.Fatalln(err)
	}

	body, err_ := ioutil.ReadAll(resp.Body)
	if err_ != nil {
		log.Fatalln(err_)
	}

	fmt.Println(string(body))

	// create a file using `os.Create` and writing to it.
	f := createFile("/home/namannimmo/Desktop/new_file.txt")
	defer closeFile(f)
	writeFile(f)

	fmt.Println("random string.")

}
