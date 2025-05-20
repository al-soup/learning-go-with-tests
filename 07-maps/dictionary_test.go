package main

import "testing"

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
		want := ErrEntryNotFound

		if err == nil {
			t.Fatal("did not receive an error as expected")
		}

		assertError(t, err, want)
	})

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "new key"
		value := "new value"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("expected to find value for key: ", key)
	}

	assertStrings(t, got, value)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
