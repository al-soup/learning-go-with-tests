package main

import (
	"testing"
	// "errors"
)

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just testing"}

		// two values are returned
		got, _ := dictionary.Search("test")
		want := "just testing"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just testing"}

		_, err := dictionary.Search("unknown")
		want := ("word not found")

		if err == nil {
			t.Fatal("did not receive an error as expected")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
