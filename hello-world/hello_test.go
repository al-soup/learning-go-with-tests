package main

import "testing"

/*
Rules for tests:
  1. It needs to be in a file with a name like xxx_test.go
	2. The test function must start with the word Test
	3. The test function takes one argument only t *testing.T
	4. To use the *testing.T type, you need to import "testing"
*/

func TestHello(t *testing.T) {
	got := Hello("Soup")
	want := "Hello, Soup"

	if got != want {
		// Placeholder %q is a string which wraps the result in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
