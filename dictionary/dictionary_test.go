package main

import (
	"testing"
)

func TestDelete(t *testing.T) {
	word := "test"
	definition := "simple test to update"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}
func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "simple test to update"
		dictionary := Dictionary{word: definition}

		newDefinition := "update in definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "simple test to update"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "newValue"
		value := "this is a test to add value"
		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "test for existing word"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExist)
	})

}

func TestSearch(t *testing.T) {
	key := "test"
	value := "this is a test"
	dictionary := Dictionary{key: value}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(key)
		assertString(t, got, value)
	})
	t.Run("unknown word", func(t *testing.T) {
		want := "word is not in dictionary"
		_, err := dictionary.Search("Unknown")

		assertError(t, err, ErrNotFound)
		assertString(t, err.Error(), want)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
