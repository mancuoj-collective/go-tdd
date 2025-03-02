package basic

import "testing"

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v but want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v but want false", got)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
