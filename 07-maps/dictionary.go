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

func (d Dictionary) Search(word string) (string, error) {
	// map lookup returns a second arg as boolean which indicates if the key was found
	entry, found := d[word]
	if !found {
		return "", errors.New("word not found")
	}

	return entry, nil
}
