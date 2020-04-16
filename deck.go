// David Walshe
// 16/04/2020

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
func (d deck) saveToDeck(fileName string) {
	err := ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// Read a deck to from a file.
func (d deck) newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cards := strings.Split(string(bs), ",")

	return cards
}

// Shuffle a decks contents
func (d deck) shuffle() {
	// Create random number generator.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// Shuffle cards.
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
