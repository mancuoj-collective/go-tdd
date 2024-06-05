package mocking

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) {
	fmt.Fprint(w, "3")
}
