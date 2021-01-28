package main

import "errors"

type Dictionary map[string]string

var dictionary = map[string]string{}

// ^ This creates an empty hash map and points dictionary at it.
// Following this approach we'll never get a runtime panic

// ~ var m map[string]string ~
//
// ^ This can cause runtime panic.

var (
	ErrWordUnavailable = errors.New("couldn't find the word you're looking for")
	ErrWordExists      = errors.New("this word already exists in Dictionary")
)

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordUnavailable
	}

	return definition, nil
}

// A map value is a pointer to a runtime.hmap structure.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordUnavailable:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) {
	d[word] = newDefinition
}

// Deletes a word in Dictionary!
func (d Dictionary) Delete(word string) {
	delete(d, word)
	// Go's builtin function delete that works on maps.
}
