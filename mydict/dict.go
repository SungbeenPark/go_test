package mydict

import "errors"

var (
	errNotFount   = errors.New("not fount")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("that word already exists")
)

//Dictionary
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFount
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFount:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFount:
		return errNotFount
	case nil:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Delete(word string) {

}
