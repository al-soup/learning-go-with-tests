package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	//fmt.Fprint takes an io.Writer (like *bytes.Buffer) and sends a string to it
	fmt.Fprint(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
