package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were lookin for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// func Search(dictionary map[string]string, word string) string {
// var wordFind string
// for key, value := range dictionary {
// 	if key == word {
// 		wordFind = value
// 	}
// }
// return dictionary[word]
// }
