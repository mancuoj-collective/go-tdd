package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"

// func(func(T) bool) === iter.Seq[T]
func countdownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := range countdownFrom(3) {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

// func main() {
// 	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
// 	Countdown(os.Stdout, sleeper)
// }
