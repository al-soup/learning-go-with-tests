package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	//fmt.Fprint takes an io.Writer (like *bytes.Buffer) and sends a string to it
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
