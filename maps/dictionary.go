package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, value string) (string, error) {
	return "", nil
}
