package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people, default in english", func(t *testing.T) {
		actual := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		actual := Hello("Jean", "French")
		expected := "Bonjour, Jean"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello to people in Chinese", func(t *testing.T) {
		actual := Hello("Li", "Chinese")
		expected := "你好, Li"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("saying hello to world when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"
		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
