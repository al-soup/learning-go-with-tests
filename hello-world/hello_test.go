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
	// t.Run is called a subtest
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Soup", "")
		want := "Hello, Soup"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Soup", "FR")
		want := "Bonjour, Soup"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Soup", "ES")
		want := "Hola, Soup"

		assertCorrectMessage(t, got, want)
	})
}

// shorthand type definition for `got` and `want`
func assertCorrectMessage(t testing.TB, got, want string) {
	// tell the test suite that this method is a helper
	t.Helper()
	if got != want {
		// Placeholder %q is a string which wraps the result in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
