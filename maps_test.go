package main

import "testing"

const (
	word = "crude"
	def  = "not yet processed or refined"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{word: def}
	// In a map, the key can only be a comparable type

	t.Run("word available", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		assertStrings(t, got, def)
	})

	t.Run("word not available", func(t *testing.T) {
		_, err := dictionary.Search("uva")
		assertDicError(t, err, ErrWordUnavailable)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		expected := "this is just a test"
		err := dictionary.Add(word, expected)

		assertDicError(t, err, nil) // This won't cause problem.
		assertDefinition(t, dictionary, word, expected)
	})

	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "another similar test")

		assertDicError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{word: def}
	newDefinition := "a new definition of crude"

	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)
}

func TestDelete(t *testing.T) {
	newWord, newDefinition := "new word", "new definition"
	dictionary := Dictionary{word: def}

	dictionary.Add(newWord, newDefinition)
	dictionary.Delete(word)
	// dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordUnavailable {
		t.Errorf("expected %q to be deleted!", word)
	}
}

func assertStrings(t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertDicError(t *testing.T, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find the definition!")
	}

	if got != definition {
		t.Errorf("got %q, expected %q", got, definition)
	}
}
