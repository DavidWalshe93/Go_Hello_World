package main

import "fmt"

// Create a new type of deck which
// is a slice of type strings.
type deck []string

// Create a receiver for deck.
func (d deck) print() {

	// Print deck contents
	for i, card := range d {
		fmt.Println(i, card)
	}
}
