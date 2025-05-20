package main

import "errors"

/*
Structs vs Maps

A struct is a composite data type that groups together variables (fields) under a single name.
A map is a collection of key-value pairs where keys must be of the same type and values must be of the same type.

Structure definition:
- Structs have fixed fields defined at compile time
- Maps can have dynamic keys added/removed at runtime

Type safety:
- Structs provide type safety for each field
- Maps have homogeneous value types (unless using interface{})
*/
type Dictionary map[string]string

var (
	ErrEntryNotFound     = errors.New("word not found")
	ErrWordAlreadyExists = errors.New("word does already exist")
)

func (d Dictionary) Search(word string) (string, error) {
	// map lookup returns a second arg as boolean which indicates if the key was found
	entry, found := d[word]
	if !found {
		return "", ErrEntryNotFound
	}

	return entry, nil
}

// Maps can be modified without passing as an address to it
// They are not reference types although they feel like it.
// When you pass a map to a func/mehtod you are copying it
// but just the pointer part and not the underlaying data structure
// So maps can be a nil value - therefor should never initialize it as such
// Rather init an empty map or use `make`
/*
	var dictionary = map[string]string{}
	// OR
	var dictionary = make(map[string]string)
*/
func (d Dictionary) Add(key, value string) error {
	// In a map already existing values for a key are overwritten
	_, err := d.Search(key)

	switch err {
	case ErrEntryNotFound:
		d[key] = value
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}
