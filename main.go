// David Walshe
// 16/04/2020

package main

import "fmt"

// Main Caller
func main() {
	// Create a slice
	cards := newDeck()

	cards.saveToFile("myDeck.txt")

	cards.shuffle()
	fmt.Println(cards)
}
