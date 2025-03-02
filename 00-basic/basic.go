package basic

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// 01 Constants
const (
	spanish            = "Spanish"
	french             = "French"
	chinese            = "Chinese"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	chineseHelloPrefix = "你好, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// 02 Integers
func Add(a, b int) int {
	return a + b
}

// 03 Iteration
func Repeat1(character string, repeatCount int) (repeated string) {
	for range repeatCount {
		repeated += character
	}
	return
}

// Strings in Go are immutable, meaning every concatenation, such as in our Repeat function,
// involves copying memory to accommodate the new string.
// This impacts performance, particularly during heavy string concatenation.
// go test -bench=.
// go test -bench=. -benchmem
func Repeat2(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func Repeat3(character string, repeatCount int) string {
	repeated := strings.Repeat(character, repeatCount)
	return repeated
}

// 04 Arrays and slices
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// 05 structs methods interfaces
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}

// not implement
// just has a method called Area and returns a float64
// any type that has this method is said to implement the Shape interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// 06 pointers errors
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// a pointer to that wallet so we can change the original value
func (w *Wallet) Deposit(amount Bitcoin) {
	// (*w).balance automatically dereferences the pointer
	// so we can just use w.balance
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// 07 maps
