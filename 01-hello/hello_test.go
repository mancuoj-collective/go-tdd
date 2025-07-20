package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Chris"
	actual := Hello("Chris")
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
