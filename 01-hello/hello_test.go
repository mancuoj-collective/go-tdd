package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world"
	actual := Hello()
	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
