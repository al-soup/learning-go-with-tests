package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// The Buffer type implements the Writer interface. Writer is a general purpose interface
	// that is used often in Go. In general Writer mean "put this data somewhere"
	// Buffer implements Writer because it has `Write(p []byte) (n int, err error)` method
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alex")

	got := buffer.String()
	want := "Hello, Alex"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
