package main

import "fmt"

// Main Caller
func main() {
	// Create a slice
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println(remainingDeck.toString())
	fmt.Println(hand.toString())
}
