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

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "new key"
		value := "new value"
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "foo"
		dictionary := Dictionary{key: "bar"}
		err := dictionary.Add(key, "baz")

		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, key, "bar")
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, errNotFound := dictionary.Search(word)
		assertError(t, errNotFound, ErrEntryNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("asdf")

		assertError(t, err, ErrWordDoesNotExist)
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
