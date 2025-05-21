package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	finalWord := "Go!"
	countdownStart := 3

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	//fmt.Fprint takes an io.Writer (like *bytes.Buffer) and sends a string to it
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
