package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		AssertEqual(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertEqual(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		AssertEqual(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		AssertEqual(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("Elodie", "Chinese")
		want := "你好, Elodie"
		AssertEqual(t, got, want)
	})
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	AssertEqual(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestRepeat1(t *testing.T) {
	repeated := Repeat1("a", 5)
	expected := "aaaaa"

	AssertEqual(t, repeated, expected)
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

// go test -cover
func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		AssertEqual(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		AssertEqual(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	AssertEqual(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	AssertEqual(t, got, want)
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
