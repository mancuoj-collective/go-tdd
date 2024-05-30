package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// 71.27 ns/op
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// 20.87 ns/op
func BenchmarkNativeRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}
}
