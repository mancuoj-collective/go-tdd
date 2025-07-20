package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Chris")
		expected := "Hello, Chris"
		if actual != expected {
			t.Errorf("actual %q, expected %q", actual, expected)
		}
	})

	t.Run("saying hello to world when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"
		if actual != expected {
			t.Errorf("actual %q, expected %q", actual, expected)
		}
	})
}
