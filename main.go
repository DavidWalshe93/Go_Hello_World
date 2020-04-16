// David Walshe
// 16/04/2020

package main

import "fmt"

// Main Caller
func main() {
	// Create a slice
	cards := newDeck()

	//hand, remainingDeck := deal(cards, 5)

	cards.saveToDeck("myDeck.txt")

	//cards.newDeckFromFile("myDeck.txt")

	//fmt.Println(cards.newDeckFromFile("myDeck.txt"))

	cards.shuffle()
	fmt.Println(cards)
}
