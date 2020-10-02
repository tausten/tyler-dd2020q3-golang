package maps

import (
	//	"errors"

	"github.com/pkg/errors"
)

type Dictionary map[string]string

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNotFound   = Error("could not find the word you were looking for")
	ErrWordExists = Error("word already exists")
)

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		err = errors.Wrap(ErrNotFound, "wrapped error")
	}
	return definition, err
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch errors.Cause(err) {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return errors.Wrap(err, "Add word failed")
	}

	return nil
}
