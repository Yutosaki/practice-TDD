package localmaps

import "errors"

var ErrNoFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string)(string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNoFound
	}
    return definition, nil
}
