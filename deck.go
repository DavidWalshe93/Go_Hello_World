package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

// Deals cards from the deck to meet the handSize passed.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Converts a deck type as a comma separated string.
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Writes a deck's contents to a file.
func (d deck) saveToDeck(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
