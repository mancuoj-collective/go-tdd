package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("Elodie", "Chinese")
		want := "你好, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertCorrectMessage(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestRepeat(t *testing.T) {
	repeated := Repeat1("a", 5)
	expected := "aaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func BenchmarkRepeat1(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	for b.Loop() {
		Repeat1("a", 5)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for b.Loop() {
		Repeat2("a", 5)
	}
}

func BenchmarkRepeat3(b *testing.B) {
	for b.Loop() {
		Repeat3("a", 5)
	}
}
