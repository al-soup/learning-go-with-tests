package main

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	// Use the writer to send the greeting to the buffer in our test

	// fmt.Printf("Hello, %s", name) <-- defaults to stdout
	// fmt.Fprintf insted uses a write to send strings to stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}
