package main

import "fmt"

// Create a new type of deck which
// is a slice of type strings.
type deck []string

// Creates a new deck
func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Create a receiver for deck.
func (d deck) print() {

	// Print deck contents
	for i, card := range d {
		fmt.Println(i, card)
	}
}
