package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "just testing"}

	got := Search(dictionary, "test")
	want := "just testing"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
